package httphpublicpresentations

import (
	"context"

	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/httphlib"
	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/views"
	"github.com/Dionid/notion-to-presentation/libs/ntp/models"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func PublicPresentationsHandlers(e *core.ServeEvent, app core.App, gctx context.Context) {
	g := e.Router.Group(httphlib.PUBLIC_PRESENTATIONS_ROUTE)

	g.GET(httphlib.PUBLIC_PRESENTATION_ROUTE, func(c echo.Context) error {
		id := c.PathParam("id")

		presentation := models.Presentation{}

		err := models.PresentationQuery(app.Dao()).
			AndWhere(dbx.HashExp{"id": id, "public": true}).
			Limit(1).
			One(&presentation)

		if err != nil {
			return err
		}

		component := views.PublicPresentationPage(&presentation)

		return component.Render(c.Request().Context(), c.Response().Writer)
	})
}
