package ntp

import (
	"fmt"
	"net/url"
)

type ReformedNotionBlock struct {
	Id             string
	Type           string // "page" | "text" | "divider" | "numbered_list" | "column" | "column_list" | "header" | "sub_header" | "code" | "image"
	Text           *string
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
				text += GetTextFormatting(part, true)

				if str, ok := part[0].(string); ok {
					text += str
				}

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

func ReformedNotionBlocks(
	domain string,
	responseChunks *NotionChunkResponse,
	block *NotionChunkResponseRecordBlock,
) ReformedNotionBlock {
	text := GetText(block.Value.Properties)
	listStartIndex := GetListStartIndex(block.Value.Format)
	var imageUrl *string

	if block.Value.Format != nil && block.Value.Format.DisplaySource != nil {
		url := fmt.Sprintf(`%s/image/%s?table=block&id=%s`, domain, url.QueryEscape(*block.Value.Format.DisplaySource), block.Value.Id)
		imageUrl = &url
	}

	rb := ReformedNotionBlock{
		Id:             block.Value.Id,
		Type:           block.Value.Type,
		Text:           text,
		ListStartIndex: listStartIndex,
		ImageUrl:       imageUrl,
		Nested:         make([]ReformedNotionBlock, len(block.Value.Content)),
	}

	for i, nestedBlockId := range block.Value.Content {
		nestedBlock := responseChunks.RecordMap.Block[nestedBlockId]
		if nestedBlock != nil {
			rb.Nested[i] = ReformedNotionBlocks(domain, responseChunks, nestedBlock)
		}
	}

	return rb
}
