package ntp

import (
	"fmt"
	"net/url"
)

type ReformedNotionBlock struct {
	Id             string
	Type           string // "page" | "text" | "divider" | "numbered_list" | "column" | "column_list" | "header" | "sub_header" | "code" | "image" | "bulleted_list" | "video" | "embed"
	Text           *string
	CodeLanguage   *string // only for type == "code"
	VideoSource    *string // only for type == "video"
	EmbedSource    *string // only for type == "embed"
	PageIcon       *string // only for type == "callout"
	ListStartIndex *int    // only for type == "numbered_list"
	ImageUrl       *string // only for type == "image"

	Nested []ReformedNotionBlock
}

func GetTextFormatting(
	part []interface{},
	open bool,
) string {
	text := ""

	if len(part) > 1 {
		if formattingPart, ok := part[1].([]interface{}); ok {
			for _, formatting := range formattingPart {
				if f, ok := formatting.([]interface{}); ok {
					for _, fff := range f {
						if ff, ok := fff.(string); ok {
							if open {
								text += fmt.Sprintf(`<%s>`, ff)
							} else {
								text += fmt.Sprintf(`</%s>`, ff)
							}
						}
					}
				}
			}
		}
	}

	return text
}

func GetText(properties *NotionChunkResponseRecordBlockValueValueProperties) *string {
	if properties != nil && len(properties.Title) > 0 {
		text := ""

		for _, part := range properties.Title {
			if part != nil && len(part) > 0 {
				content, ok := part[0].(string)
				if !ok {
					continue
				}

				text += GetTextFormatting(part, true)

				text += content

				text += GetTextFormatting(part, false)
			}
		}

		return &text
	}
	return nil
}

func GetListStartIndex(format *NotionChunkResponseRecordBlockValueValueFormat) *int {
	if format != nil {
		return format.ListStartIndex
	}
	return nil
}

func GetImageUrl(format *NotionChunkResponseRecordBlockValueValueFormat) *string {
	if format != nil {
		return format.DisplaySource
	}
	return nil
}

func GetCodeLanguage(properties *NotionChunkResponseRecordBlockValueValueProperties) *string {
	if properties != nil && len(properties.Language) > 0 {
		if language, ok := properties.Language[0][0].(string); ok {
			return &language
		}
	}
	return nil
}

func GetEmbedSource(properties *NotionChunkResponseRecordBlockValueValueProperties) *string {
	if properties != nil && len(properties.Source) > 0 {
		if source, ok := properties.Source[0][0].(string); ok {
			return &source
		}
	}
	return nil
}

func ReformedNotionBlocks(
	domain string,
	responseChunks *NotionChunkResponse,
	block *NotionChunkResponseRecordBlock,
) ReformedNotionBlock {
	text := GetText(block.Value.Properties)
	listStartIndex := GetListStartIndex(block.Value.Format)

	rb := ReformedNotionBlock{
		Id:             block.Value.Id,
		Type:           block.Value.Type,
		Text:           text,
		ListStartIndex: listStartIndex,
		Nested:         make([]ReformedNotionBlock, len(block.Value.Content)),
	}

	if rb.Id == "2ce12a9b-babd-4813-8f6d-d7878ad70efb" {
		fmt.Println("!!! lorem ipsum", block.Value.Type)
	}

	if rb.Type == "callout" {
		if block.Value.Format.PageIcon != nil {
			rb.PageIcon = block.Value.Format.PageIcon
		}
	}

	if block.Value.Format != nil && block.Value.Format.DisplaySource != nil && block.Value.Type == "image" {
		url := fmt.Sprintf(`%s/image/%s?table=block&id=%s`, domain, url.QueryEscape(*block.Value.Format.DisplaySource), block.Value.Id)
		rb.ImageUrl = &url
	}

	if block.Value.Type == "code" {
		rb.CodeLanguage = GetCodeLanguage(block.Value.Properties)
	}

	if block.Value.Type == "video" {
		rb.VideoSource = block.Value.Format.DisplaySource
	}

	if block.Value.Type == "embed" {
		rb.EmbedSource = block.Value.Format.DisplaySource
	}

	for i, nestedBlockId := range block.Value.Content {
		nestedBlock := responseChunks.RecordMap.Block[nestedBlockId]
		if nestedBlock != nil {
			rb.Nested[i] = ReformedNotionBlocks(domain, responseChunks, nestedBlock)
		}
	}

	return rb
}
