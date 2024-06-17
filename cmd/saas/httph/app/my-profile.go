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

func MyProfileHandlers(pg *echo.Group, app core.App, gctx context.Context) {
	g := pg.Group(httphlib.APP_MY_PROFILE_ROUTE)

	g.GET("", func(c echo.Context) error {
		userRecord, _ := httphlib.GetAuthedUserRecord(c)
		if userRecord == nil {
			return c.Redirect(302, httphlib.SIGN_IN_ROUTE)
		}

		user := models.User{}

		err := app.Dao().DB().
			Select("id", "email", "name", "description").
			From("users").
			AndWhere(dbx.Like("id", userRecord.Id)).
			Limit(1).
			One(&user)
		if err != nil {
			return err
		}

		component := views.MyProfilePage(&user)

		return component.Render(c.Request().Context(), c.Response().Writer)
	})
}
