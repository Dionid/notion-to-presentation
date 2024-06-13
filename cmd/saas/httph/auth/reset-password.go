package httphauth

import (
	"context"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/httphlib"
	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/views"
)

func ResetPasswordHandlers(e *core.ServeEvent, app core.App, gctx context.Context) {
	e.Router.GET(httphlib.RESET_PASSWORD_ROUTE, func(c echo.Context) error {
		// # If already logged in, redirect to home page
		err := httphlib.RedirectIfAuthorized(c, httphlib.APP_ROUTE)
		if err != nil {
			return err
		}

		component := views.ResetPassword("")

		return component.Render(gctx, c.Response().Writer)
	}, apis.ActivityLogger(app), apis.RequireGuestOnly())

	return
}
