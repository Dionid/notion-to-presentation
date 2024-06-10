package httphpreview

import (
	"context"
	"fmt"
	"net/url"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/httphlib"
	"github.com/Dionid/notion-to-presentation/libs/ntp"
	"github.com/Dionid/notion-to-presentation/libs/ntp/models"
)

func PreviewHandlers(e *core.ServeEvent, app core.App, gctx context.Context, previewId string) {
	e.Router.GET(httphlib.PREVIEW_ROUTE, func(c echo.Context) error {
		data := struct {
			Url string `query:"url" json:"url"`
		}{}
		if err := c.Bind(&data); err != nil {
			return apis.NewBadRequestError("Failed to read request data", err)
		}

		if data.Url == "https://it-kachalka.notion.site/28-May-2024-5c488ae7df5249c2ae0639a20fff32ef" {
			presentation := &models.Presentation{}

			err := models.PresentationQuery(app.Dao()).
				AndWhere(dbx.HashExp{"id": previewId}).
				Limit(1).
				One(presentation)

			if err != nil {
				return err
			}

			html := presentation.Html

			html += `
			<style>
				` + presentation.CustomCss + `
			</style>
			`

			html += `
			<script>
				(() => {
					const revealPresentation = new Reveal({
						hash: true,
						plugins: [RevealMarkdown, RevealHighlight, RevealNotes],
						embedded: true,
					});
				
					revealPresentation.initialize();
				})()
			</script>
			`

			return c.HTML(200, html)
		}

		// # Form urls
		parsedUrl, err := url.Parse(data.Url)
		if err != nil {
			return fmt.Errorf("Failed to parse url: %w", err)
		}

		domain := fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host)
		mainPageId := ntp.ExtractPageIdFromUrl(parsedUrl)

		// # Get page blocks
		responseChunks, err := ntp.GetNotionBlocks(domain, mainPageId)
		if err != nil {
			return fmt.Errorf("Failed to get notion blocks: %w", err)
		}

		chunkedBlocks, err := ntp.FormChunkedBlocks(domain, responseChunks, mainPageId)
		if err != nil {
			return fmt.Errorf("Failed to form chunked blocks: %w", err)
		}

		// # Form html
		html := ntp.FormRevealHtml(chunkedBlocks)

		html += `
<script>
	(() => {
		const revealPresentation = new Reveal({
			hash: true,
			plugins: [RevealMarkdown, RevealHighlight, RevealNotes],
			embedded: true,
		});
	
		revealPresentation.initialize();
	})()
</script>
`

		return c.HTML(200, html)
	}, apis.ActivityLogger(app), apis.RequireGuestOnly())

	return
}
