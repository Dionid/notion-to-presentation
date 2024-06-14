package ntp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type NotionChunkResponseRecordBlockValueValueProperties struct {
	Title    [][]interface{} `json:"title,omitempty"`
	Language [][]interface{} `json:"language,omitempty"`
	Source   [][]interface{} `json:"source,omitempty"`
}

type NotionChunkResponseRecordBlockValueValueFormat struct {
	ListStartIndex *int    `json:"list_start_index,omitempty"`
	DisplaySource  *string `json:"display_source,omitempty"`
	PageIcon       *string `json:"page_icon,omitempty"`
}

type NotionChunkResponseRecordBlockValueValue struct {
	Id         string                                              `json:"id"`
	Content    []string                                            `json:"content"`
	Type       string                                              `json:"type"` // "page" | "text" | "divider" | "numbered_list"
	Properties *NotionChunkResponseRecordBlockValueValueProperties `json:"properties,omitempty"`
	Format     *NotionChunkResponseRecordBlockValueValueFormat     `json:"format,omitempty"`
}

type NotionChunkResponseRecordBlockValue struct {
	Role  string                                   `json:"role"`
	Value NotionChunkResponseRecordBlockValueValue `json:"value"`
}

type NotionChunkResponseRecordBlock struct {
	SpaceId *string                                  `json:"spaceId,omitempty"`
	Value   NotionChunkResponseRecordBlockValueValue `json:"value"`
}

type NotionChunkResponseRecord struct {
	Block map[string]*NotionChunkResponseRecordBlock `json:"block"`
}

type NotionChunkResponseCursorStackElement struct {
	Index int    `json:"index"`
	Id    string `json:"id"`
	Table string `json:"table"`
}

type NotionChunkResponseCursor struct {
	Stack [][]NotionChunkResponseCursorStackElement `json:"stack"`
}

type NotionChunkResponse struct {
	Cursor    *NotionChunkResponseCursor `json:"cursor"`
	RecordMap NotionChunkResponseRecord  `json:"recordMap"`
}

func ExtractPageIdFromUrl(parsedUrl *url.URL) string {
	splittedPath := strings.Split(parsedUrl.Path, "-")
	originalId := strings.ReplaceAll(splittedPath[len(splittedPath)-1], "/", "")
	mainPageId := fmt.Sprintf("%s-%s-%s-%s-%s", originalId[0:8], originalId[8:12], originalId[12:16], originalId[16:20], originalId[20:])

	return mainPageId
}

type LoadCachedPageChunkRequestBodyPage struct {
	Id string `json:"id"`
	// SpaceId string `json:"spaceId"`
}

type LoadCachedPageChunkRequestBody struct {
	Page            LoadCachedPageChunkRequestBodyPage `json:"page"`
	Limit           int                                `json:"limit"`
	Cursor          NotionChunkResponseCursor          `json:"cursor"`
	ChunkNumber     int                                `json:"chunkNumber"`
	VerticalColumns bool                               `json:"verticalColumns"`
}

func GetNotionBlocksRequest(
	logger Logger,
	domain string,
	mainPageId string,
	requestBody *LoadCachedPageChunkRequestBody,
) (*NotionChunkResponse, error) {
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(
		fmt.Sprintf("%s/api/v3/loadCachedPageChunk", domain),
		"application/json",
		bytes.NewBuffer(requestBodyBytes),
	)

	if err != nil {
		return nil, fmt.Errorf("Error sending request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		logger.Debug("Error from Notion response, body:", body)
		return nil, errors.New(fmt.Sprintf("status code error: %d", res.StatusCode))
	}

	responseChunks := NotionChunkResponse{}

	err = json.Unmarshal(body, &responseChunks)
	if err != nil {
		return nil, err
	}

	return &responseChunks, nil
}

func GetNotionBlocksRecursive(
	logger Logger,
	domain string,
	mainPageId string,
	stack [][]NotionChunkResponseCursorStackElement,
	chunkNumber int,
) (*NotionChunkResponse, error) {
	requestBody := LoadCachedPageChunkRequestBody{
		Page: LoadCachedPageChunkRequestBodyPage{
			Id: mainPageId,
		},
		Limit: 100,
		Cursor: NotionChunkResponseCursor{
			Stack: stack,
		},
		ChunkNumber:     chunkNumber,
		VerticalColumns: false,
	}

	responseChunks, err := GetNotionBlocksRequest(logger, domain, mainPageId, &requestBody)
	if err != nil {
		return nil, err
	}

	if responseChunks.Cursor.Stack != nil && len(responseChunks.Cursor.Stack) > 0 {
		nestedResponseChunks, err := GetNotionBlocksRecursive(logger, domain, mainPageId, responseChunks.Cursor.Stack, chunkNumber+1)
		if err != nil {
			return nil, err
		}

		for blockId, block := range nestedResponseChunks.RecordMap.Block {
			responseChunks.RecordMap.Block[blockId] = block
		}
	}

	return responseChunks, nil
}

type Logger interface {
	Error(msg string, args ...any)
	Debug(msg string, args ...any)
}

func GetNotionBlocks(
	logger Logger,
	domain string,
	mainPageId string,
) (*NotionChunkResponse, error) {
	responseChunks, err := GetNotionBlocksRecursive(logger, domain, mainPageId, [][]NotionChunkResponseCursorStackElement{}, 0)
	if err != nil {
		return nil, err
	}

	return responseChunks, nil
}

func ExtractPageTitle(responseChunks *NotionChunkResponse, mainPageId string) (string, error) {
	mainPageBlock := responseChunks.RecordMap.Block[mainPageId]

	if mainPageBlock.Value.Id != mainPageId {
		return "", errors.New("Page not found")
	}

	pageTitle, ok := mainPageBlock.Value.Properties.Title[0][0].(string)
	if !ok {
		return "", errors.New("Page title not found")
	}

	return pageTitle, nil
}

func FormChunkedBlocks(logger Logger, domain string, responseChunks *NotionChunkResponse, mainPageId string) ([][]ReformedNotionBlock, error) {
	mainPageBlock := responseChunks.RecordMap.Block[mainPageId]

	if mainPageBlock.Value.Id != mainPageId {
		return nil, errors.New("Page not found")
	}

	// # Sort blocks
	sortedBlocks := make([]NotionChunkResponseRecordBlock, len(mainPageBlock.Value.Content))

	for i, blockId := range mainPageBlock.Value.Content {
		block := responseChunks.RecordMap.Block[blockId]
		if block == nil {
			// TODO: Think more about it
			logger.Error("Block not found", blockId)
			continue
			// return nil, errors.New("Block not found " + blockId)
		}

		sortedBlocks[i] = *block
	}

	// # Form nested blocks
	var reformedBlocks []ReformedNotionBlock

	for _, block := range sortedBlocks {
		rb := ReformedNotionBlocks(domain, responseChunks, &block)

		reformedBlocks = append(reformedBlocks, rb)
	}

	// // # Chunk by divider
	var chunkedBlocks [][]ReformedNotionBlock
	chunkedBlocks = append(chunkedBlocks, []ReformedNotionBlock{})

	for _, block := range reformedBlocks {
		if block.Type == "divider" {
			chunkedBlocks = append(chunkedBlocks, []ReformedNotionBlock{})
		} else {
			chunkedBlocks[len(chunkedBlocks)-1] = append(chunkedBlocks[len(chunkedBlocks)-1], block)
		}
	}

	return chunkedBlocks, nil
}
