package ntp

import (
	"fmt"
)

func FormSectionContent(
	blocks []ReformedNotionBlock,
) string {
	html := ""
	isNumbered := false

	for _, block := range blocks {
		// fmt.Println("Block type: %s", block.Type)

		if isNumbered && block.Type != "numbered_list" {
			html += "</ol>"
			isNumbered = false
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
		case "text":
			if block.Text != nil {
				html += "<p>" + *block.Text + "</p>"
			}
		case "code":
			if block.Text != nil {
				html += "<pre><code>" + *block.Text + "</code></pre>"
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
		case "column":
			html += "<div class=\"column\">"
		case "column_list":
			html += `<div class="columns-holder">`
		case "image":
			if block.ImageUrl != nil {
				html += fmt.Sprintf(`<img src="%s" alt="image" />`, *block.ImageUrl)
			}
		case "page":
			break
		default:
			fmt.Println("Block type: ", block.Type)
		}

		html += "\n"

		// ## Form nested blocks
		if len(block.Nested) > 0 {
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
