package httphauth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tokens"
	"github.com/pocketbase/pocketbase/tools/security"
	"golang.org/x/crypto/bcrypt"

	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/httphlib"
	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/views"
)

func SignInHandlers(e *core.ServeEvent, app core.App, gctx context.Context) {
	e.Router.GET(httphlib.SIGN_IN_ROUTE, func(c echo.Context) error {
		// # If already logged in, redirect to home page
		err := httphlib.RedirectIfAuthorized(c, httphlib.APP_ROUTE)
		if err != nil {
			return err
		}

		component := views.SignInPage("")

		return component.Render(gctx, c.Response().Writer)
	}, apis.ActivityLogger(app), apis.RequireGuestOnly())

	e.Router.POST(httphlib.SIGN_IN_ROUTE, func(c echo.Context) error {
		// # If already logged in, redirect to home page
		err := httphlib.RedirectIfAuthorized(c, httphlib.APP_ROUTE)
		if err != nil {
			return err
		}

		data := struct {
			Email    string `form:"email" json:"email"`
			Password string `form:"password" json:"password"`
		}{}
		if err := c.Bind(&data); err != nil {
			return apis.NewBadRequestError("Failed to read request data", err)
		}

		if data.Email == "" || data.Password == "" {
			component := views.SignInPageForm("Email and Password are required")

			return component.Render(gctx, c.Response().Writer)
		}

		user, err := app.Dao().FindAuthRecordByEmail("users", data.Email)

		if err != nil {
			app.Logger().Error(fmt.Sprintf("Failed to fetch user: %v", err))

			component := views.SignInPageForm("No user")

			return component.Render(gctx, c.Response().Writer)
		}

		app.Logger().Debug(fmt.Sprintf("User: %+v", user))

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash()), []byte(data.Password))
		if err != nil {
			component := views.SignInPageForm("Invalid email or password")

			return component.Render(gctx, c.Response().Writer)
		}

		jwt, err := security.NewJWT(
			jwt.MapClaims{
				"id":           user.Id,
				"type":         tokens.TypeAuthRecord,
				"collectionId": user.Collection().Id,
				"email":        user.Email(),
			},
			(user.TokenKey() + app.Settings().RecordAuthToken.Secret),
			app.Settings().RecordAuthToken.Duration,
		)

		httphlib.SetCookie(c, jwt)

		c.Response().Header().Set("HX-Redirect", httphlib.APP_ROUTE)

		return c.NoContent(http.StatusOK)
	}, apis.ActivityLogger(app), apis.RequireGuestOnly())

	return
}
