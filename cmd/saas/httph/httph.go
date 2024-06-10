package httph

import (
	"context"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	httphapp "github.com/Dionid/notion-to-presentation/cmd/saas/httph/app"
	httphauth "github.com/Dionid/notion-to-presentation/cmd/saas/httph/auth"
	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/httphlib"
	httphpreview "github.com/Dionid/notion-to-presentation/cmd/saas/httph/preview"
	httphpublicpresentations "github.com/Dionid/notion-to-presentation/cmd/saas/httph/public-presentations"
	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/views"
)

type Config struct {
	PreviewId string
}

func InitApi(config Config, app core.App, gctx context.Context) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// # Static
		// e.Router.GET("/public/*", apis.StaticDirectoryHandler(os.DirFS("public"), true))

		e.Router.Static("/public", "public")

		// # PB Auth
		e.Router.Use(httphlib.LoadAuthContextFromCookieMiddleware(app))

		e.Router.Use(apis.ActivityLogger(app))

		// # Landing
		e.Router.GET("/", func(c echo.Context) error {
			err := httphlib.RedirectIfAuthorized(c, httphlib.APP_ROUTE)
			if err != nil {
				return err
			}

			component := views.IndexPage()

			return component.Render(gctx, c.Response().Writer)
		})

		// # Auth
		// ## Sign in
		httphauth.SignInHandlers(e, app, gctx)
		// ## Sign up
		httphauth.SignUpHandlers(e, app, gctx)

		// # Preview
		httphpreview.PreviewHandlers(e, app, gctx, config.PreviewId)

		// # App
		httphapp.AppHandlers(e, app, gctx)

		// # Public Presentations
		httphpublicpresentations.PublicPresentationsHandlers(e, app, gctx)

		return nil
	})
}
