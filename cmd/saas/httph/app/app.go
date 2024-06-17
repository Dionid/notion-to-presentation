package httphapp

import (
	"context"

	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/httphlib"
	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/views"
	"github.com/Dionid/notion-to-presentation/libs/ntp/models"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func AppHandlers(e *core.ServeEvent, app core.App, gctx context.Context) {
	g := e.Router.Group(httphlib.APP_ROUTE)

	g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, err := httphlib.GetAuthedUserRecordOrDeleteSession(c)
			if err != nil {
				return err
			}

			return next(c)
		}
	})

	g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "no-store")
			return next(c)
		}
	})

	g.GET("", func(c echo.Context) error {
		userRecord, _ := httphlib.GetAuthedUserRecord(c)
		if userRecord == nil {
			return c.Redirect(302, httphlib.SIGN_IN_ROUTE)
		}

		presentations := []*models.Presentation{}

		err := models.PresentationQuery(app.Dao()).
			AndWhere(dbx.HashExp{"user_id": userRecord.Id}).
			OrderBy("created desc").
			All(&presentations)
		if err != nil {
			return err
		}

		component := views.AppIndexPage(presentations)

		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	g.GET(httphlib.APP_PRESENTATION_ROUTE, func(c echo.Context) error {
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

		component := views.PresentationPage(&presentation)

		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	ResyncPresentationHandlers(g, app, gctx)
	PresentationHandlers(g, app, gctx)
	MyProfileHandlers(g, app, gctx)
}
