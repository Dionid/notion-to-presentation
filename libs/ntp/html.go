package ntp

import (
	"fmt"
	"strings"
)

func FormSectionContent(
	blocks []ReformedNotionBlock,
) string {
	html := ""
	isNumbered := false
	isBulleted := false

	for _, block := range blocks {
		// fmt.Println("Block type: %s", block.Type)

		if isNumbered && block.Type != "numbered_list" {
			html += "</ol>"
			isNumbered = false
		}

		if isBulleted && block.Type != "bulleted_list" {
			html += "</ul>"
			isBulleted = false
		}

		switch block.Type {
		case "header":
			if block.Text != nil {
				html += "<h1>" + *block.Text + "</h1>"
			}
		case "sub_header":
			if block.Text != nil {
				html += "<h2>" + *block.Text + "</h2>"
			}
		case "sub_sub_header":
			if block.Text != nil {
				html += "<h3>" + *block.Text + "</h3>"
			}
		case "text":
			if block.Text != nil {
				html += "<p>" + *block.Text + "</p>"
			}
		case "code":
			if block.Text != nil {
				html += fmt.Sprintf("<pre><code class='language-%s'>%s</code></pre>", strings.ToLower(*block.CodeLanguage), *block.Text)
			}
		case "video":
			if block.VideoSource != nil {
				if strings.Contains(*block.VideoSource, "youtube") {
					html += fmt.Sprintf(`<iframe class="w-full" style="height: 60vh" src="%s" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>`, *block.VideoSource)
					break
				}

				html += fmt.Sprintf(`<video src="%s" controls class="w-full"></video>`, *block.VideoSource)
			}
		case "embed":
			if block.EmbedSource != nil {
				html += fmt.Sprintf(`<iframe src="%s" class="w-full" style="height: 60vh"></iframe>`, *block.EmbedSource)
			}
		case "numbered_list":
			if !isNumbered {
				start := 1

				if block.ListStartIndex != nil {
					start = *block.ListStartIndex
				}

				html += fmt.Sprintf(`<ol start="%d">`, start)
				isNumbered = true
			}
			if block.Text != nil {
				html += "<li>" + *block.Text + "</li>"
			}
		case "bulleted_list":
			if !isBulleted {
				start := 1

				if block.ListStartIndex != nil {
					start = *block.ListStartIndex
				}

				html += fmt.Sprintf(`<ul start="%d">`, start)
				isBulleted = true
			}
			if block.Text != nil {
				html += "<li>" + *block.Text + "</li>"
			}
		case "column":
			html += "<div class=\"column\">"
		case "column_list":
			html += `<div class="columns-holder">`
		case "image":
			if block.ImageUrl != nil {
				html += fmt.Sprintf(`<img src="%s" alt="image" />`, *block.ImageUrl)
			}
		// case "audio":
		// 	if block.AudioSource != nil {
		// 		html += fmt.Sprintf(`<audio src="%s" controls></audio>`, *block.AudioSource)
		// 	}
		// case "file":
		// 	if block.FileSource != nil {
		// 		html += fmt.Sprintf(`<a href="%s" target="_blank" download>%s</a>`, *block.FileSource, *block.Text)
		// 	}
		case "callout":
			if block.Text != nil {
				innerText := ""

				for _, nestedBlock := range block.Nested {
					if nestedBlock.Text != nil {
						innerText += "<br><br>"
						innerText += *nestedBlock.Text
					}
				}

				html += fmt.Sprintf(`
				<div class="card" style="background-color: #383838;">
					<div class="card-body flex-row gap-4">
					 <div>
					 	%s
					 </div>
					 <div>
					 	%s
						%s
					 </div>
					</div>
				</div>
				`, *block.PageIcon, *block.Text, innerText)
			}
		case "quote":
			if block.Text != nil {
				innerText := ""

				for _, nestedBlock := range block.Nested {
					if nestedBlock.Text != nil {
						innerText += "<br><br>"
						innerText += *nestedBlock.Text
					}
				}

				html += fmt.Sprintf(`
				<div style="border-left: 4px solid; padding-left: 30px;">
						%s
						%s
				</div>
				`, *block.Text, innerText)
			}
		case "audio":
		case "file":
		case "toggle":
		case "page":
			break
		default:
			fmt.Println("Unknown block type: ", block.Type)
		}

		html += "\n"

		// ## Form nested blocks
		if len(block.Nested) > 0 && block.Type != "callout" && block.Type != "quote" {
			html += FormSectionContent(block.Nested)
		}

		// # After nested blocks
		if block.Type == "column" {
			html += "</div>"
		}

		if block.Type == "column_list" {
			html += "</div>"
		}
	}

	if isNumbered {
		html += "</ol>"
	}

	if isBulleted {
		html += "</ul>"
	}

	return html
}

func FormRevealHtml(
	chunkedBlocks [][]ReformedNotionBlock,
) string {
	html := "<div class=\"reveal\">"
	html += "<div class=\"slides\" style=\"text-align: left\">"

	for _, chunk := range chunkedBlocks {
		html += "<section>"

		html += FormSectionContent(chunk)

		html += "</section>"
	}

	html += "</div></div>"

	return html
}

func FormFullHtmlPage(
	pageTitle string,
	chunkedBlocks [][]ReformedNotionBlock,
) string {
	html := fmt.Sprintf(`
	<html>
	<head>
	<title>%s</title>
	<link rel="stylesheet" href="dist/reset.css" />
    <link rel="stylesheet" href="dist/reveal.css" />
    <link rel="stylesheet" href="dist/theme/black.css" />
	<link rel="stylesheet" href="plugin/highlight/monokai.css" />
	<link rel="stylesheet" href="custom.css" />
	</head>
	<body>`, pageTitle)

	html += FormRevealHtml(chunkedBlocks)

	html += `<script src="dist/reveal.js"></script>
    <script src="plugin/notes/notes.js"></script>
    <script src="plugin/markdown/markdown.js"></script>
    <script src="plugin/highlight/highlight.js"></script>
    <script>
      Reveal.initialize({
        hash: true,
        plugins: [RevealMarkdown, RevealHighlight, RevealNotes],
      });
    </script>
	`

	html += `</body></html>`

	return html
}
