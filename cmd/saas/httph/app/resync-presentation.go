package httphapp

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/httphlib"
	"github.com/Dionid/notion-to-presentation/libs/ntp"
	"github.com/Dionid/notion-to-presentation/libs/ntp/models"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func ResyncPresentationHandlers(g *echo.Group, app core.App, gctx context.Context) {
	g.GET(httphlib.APP_PRESENTATION_RESYNC_ROUTE, func(c echo.Context) error {
		userRecord, _ := httphlib.GetAuthedUserRecord(c)
		if userRecord == nil {
			return c.Redirect(302, httphlib.SIGN_IN_ROUTE)
		}

		id := c.PathParam("id")

		presentation := models.Presentation{}

		err := models.PresentationQuery(app.Dao()).
			AndWhere(dbx.HashExp{"id": id, "user_id": userRecord.Id}).
			Limit(1).
			One(&presentation)

		if err != nil {
			return err
		}

		parsedUrl, err := url.Parse(presentation.NotionPageUrl)
		if err != nil {
			return err
		}

		domain := fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host)
		mainPageId := ntp.ExtractPageIdFromUrl(parsedUrl)

		// # Get page blocks
		responseChunks, err := ntp.GetNotionBlocks(domain, mainPageId)
		if err != nil {
			return err
		}

		chunkedBlocks, err := ntp.FormChunkedBlocks(domain, responseChunks, mainPageId)
		if err != nil {
			return err
		}

		// # Form html
		html := ntp.FormRevealHtml(chunkedBlocks)

		return c.JSON(200, map[string]string{"result": html})
	})
}
