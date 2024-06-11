package ntp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type NotionChunkResponseRecordBlockValueValueProperties struct {
	Title    [][]interface{} `json:"title,omitempty"`
	Language [][]interface{} `json:"language,omitempty"`
}

type NotionChunkResponseRecordBlockValueValueFormat struct {
	ListStartIndex *int    `json:"list_start_index,omitempty"`
	DisplaySource  *string `json:"display_source,omitempty"`
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
	SpaceId string                                   `json:"spaceId"`
	Value   NotionChunkResponseRecordBlockValueValue `json:"value"`
}

type NotionChunkResponseRecord struct {
	Block map[string]*NotionChunkResponseRecordBlock `json:"block"`
}

type NotionChunkResponse struct {
	Cursor    map[string]interface{}    `json:"cursor"`
	RecordMap NotionChunkResponseRecord `json:"recordMap"`
}

func ExtractPageIdFromUrl(parsedUrl *url.URL) string {
	splittedPath := strings.Split(parsedUrl.Path, "-")
	originalId := strings.ReplaceAll(splittedPath[len(splittedPath)-1], "/", "")
	mainPageId := fmt.Sprintf("%s-%s-%s-%s-%s", originalId[0:8], originalId[8:12], originalId[12:16], originalId[16:20], originalId[20:])

	return mainPageId
}

func GetNotionBlocks(domain string, mainPageId string) (*NotionChunkResponse, error) {
	res, err := http.Post(
		fmt.Sprintf("%s/api/v3/loadCachedPageChunk", domain),
		"application/json",
		bytes.NewBuffer([]byte(
			fmt.Sprintf(`{"page":{"id":"%s"},"limit":100,"cursor":{"stack":[]},"chunkNumber":0,"verticalColumns":false}`, mainPageId),
		)),
	)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code error: %d", res.StatusCode))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return nil, err
	}

	responseChunks := NotionChunkResponse{}

	err = json.Unmarshal(body, &responseChunks)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &responseChunks, nil
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

func FormChunkedBlocks(domain string, responseChunks *NotionChunkResponse, mainPageId string) ([][]ReformedNotionBlock, error) {
	mainPageBlock := responseChunks.RecordMap.Block[mainPageId]

	if mainPageBlock.Value.Id != mainPageId {
		return nil, errors.New("Page not found")
	}

	// # Sort blocks
	sortedBlocks := make([]NotionChunkResponseRecordBlock, len(mainPageBlock.Value.Content))

	for i, blockId := range mainPageBlock.Value.Content {
		block := responseChunks.RecordMap.Block[blockId]
		if block == nil {
			return nil, errors.New("Block not found")
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
