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
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

func PresentationHandlers(g *echo.Group, app core.App, gctx context.Context) {
	g.POST(httphlib.APP_PRESENTATION_NEW_ROUTE, func(c echo.Context) error {
		userRecord, _ := httphlib.GetAuthedUserRecord(c)
		if userRecord == nil {
			return c.Redirect(302, httphlib.SIGN_IN_ROUTE)
		}

		count := struct {
			Count int `db:"count"`
		}{}

		err := app.
			Dao().
			DB().
			Select("count(*) as count").
			From("presentation").
			AndWhere(dbx.HashExp{"user_id": userRecord.Id}).
			One(&count)

		if err != nil {
			return err
		}

		if count.Count >= 10 {
			return apis.NewBadRequestError("You have reached the limit of 10 presentations", nil)
		}

		data := struct {
			NotionUrl string `json:"notionUrl"`
		}{}
		if err := c.Bind(&data); err != nil {
			return apis.NewBadRequestError("Failed to read request data", err)
		}

		parsedUrl, err := url.Parse(data.NotionUrl)
		if err != nil {
			return err
		}

		domain := fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host)
		mainPageId := ntp.ExtractPageIdFromUrl(parsedUrl)

		// # Get page blocks
		responseChunks, err := ntp.GetNotionBlocks(app.Logger(), domain, mainPageId)
		if err != nil {
			return err
		}

		title := ""

		if responseChunks.RecordMap.Block[mainPageId] != nil {
			if responseChunks.RecordMap.Block[mainPageId].Value.Properties != nil {
				if responseChunks.RecordMap.Block[mainPageId].Value.Properties.Title != nil {
					if responseChunks.RecordMap.Block[mainPageId].Value.Properties.Title[0] != nil {
						if responseChunks.RecordMap.Block[mainPageId].Value.Properties.Title[0][0] != nil {
							title = responseChunks.RecordMap.Block[mainPageId].Value.Properties.Title[0][0].(string)
						}
					}
				}
			}
		}

		chunkedBlocks, err := ntp.FormChunkedBlocks(app.Logger(), domain, responseChunks, mainPageId)
		if err != nil {
			return err
		}

		// # Form html
		html := ntp.FormRevealHtml(chunkedBlocks)

		newPresentation := models.Presentation{
			NotionPageUrl:  data.NotionUrl,
			Html:           html,
			CustomCss:      "",
			UserID:         userRecord.Id,
			Title:          title,
			Description:    "",
			Public:         false,
			Theme:          "night",
			Customizations: types.JsonMap{},
		}

		if err := app.Dao().Save(&newPresentation); err != nil {
			return err
		}

		return c.JSON(200, struct {
			Id string `json:"id"`
		}{
			Id: newPresentation.Id,
		})
	})
}
