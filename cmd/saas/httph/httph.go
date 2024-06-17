package httph

import (
	"context"
	"embed"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	httphapp "github.com/Dionid/notion-to-presentation/cmd/saas/httph/app"
	httphauth "github.com/Dionid/notion-to-presentation/cmd/saas/httph/auth"
	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/httphlib"
	httphpreview "github.com/Dionid/notion-to-presentation/cmd/saas/httph/preview"
	httphpublicpresentations "github.com/Dionid/notion-to-presentation/cmd/saas/httph/public-presentations"
	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/views"
	"github.com/Dionid/notion-to-presentation/libs/file"
)

type Config struct {
	Env       string
	PreviewId string
}

//go:embed public
var publicAssets embed.FS

func InitApi(config Config, app core.App, gctx context.Context) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// # Body limit
		e.Router.Use(middleware.BodyLimit(2 * 1024 * 1024))

		// # Static
		if config.Env == "PRODUCTION" {
			file.CopyFromEmbed(publicAssets, "public", "./public")
			e.Router.Static("/public", "./public")
		} else {
			e.Router.Static("/public", "./httph/public")
		}

		// e.Router.Use(
		// 	middleware.StaticWithConfig(
		// 		middleware.StaticConfig{
		// 			Root:       "",
		// 			Browse:     false,
		// 			Filesystem: publicAssets,
		// 		},
		// 	),
		// )

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

			return component.Render(c.Request().Context(), c.Response().Writer)
		})

		// # Auth
		// ## Sign in
		httphauth.SignInHandlers(e, app, gctx)
		// ## Sign up
		httphauth.SignUpHandlers(e, app, gctx)
		// ## Reset password
		httphauth.ResetPasswordHandlers(e, app, gctx)

		// # Preview
		httphpreview.PreviewHandlers(e, app, gctx, config.PreviewId)

		// # App
		httphapp.AppHandlers(e, app, gctx)

		// # Public Presentations
		httphpublicpresentations.PublicPresentationsHandlers(e, app, gctx)

		return nil
	})
}
