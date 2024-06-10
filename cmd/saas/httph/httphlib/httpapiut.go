package httphlib

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tokens"
	"github.com/pocketbase/pocketbase/tools/security"
	"github.com/spf13/cast"
)

const AUTH_COOKIE_NAME = "pb_auth"

func LoadAuthContextFromCookieMiddleware(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenCookie, err := GetAuthCookie(c)
			if err != nil || tokenCookie.Value == "" {
				return next(c) // no token cookie
			}

			token := tokenCookie.Value

			claims, _ := security.ParseUnverifiedJWT(token)
			tokenType := cast.ToString(claims["type"])

			switch tokenType {
			case tokens.TypeAdmin:
				admin, err := app.Dao().FindAdminByToken(
					token,
					app.Settings().AdminAuthToken.Secret,
				)
				if err == nil && admin != nil {
					// "authenticate" the admin
					c.Set(apis.ContextAdminKey, admin)
				}
			case tokens.TypeAuthRecord:
				record, err := app.Dao().FindAuthRecordByToken(
					token,
					app.Settings().RecordAuthToken.Secret,
				)
				if err == nil && record != nil {
					// "authenticate" the app user
					c.Set(apis.ContextAuthRecordKey, record)
				}
			}

			return next(c)
		}
	}
}

func GetAuthedUserRecord(c echo.Context) (*models.Record, bool) {
	record, ok := c.Get(apis.ContextAuthRecordKey).(*models.Record)

	return record, ok
}

func GetAuthedUserRecordOrDeleteSession(c echo.Context) (*models.Record, error) {
	record, ok := GetAuthedUserRecord(c)
	if record == nil || !ok {
		DeleteAuthCookie(c)

		return nil, c.Redirect(302, SIGN_IN_ROUTE)
	}

	return record, nil
}

func DeleteAuthCookie(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = AUTH_COOKIE_NAME
	cookie.Value = ""
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	c.SetCookie(cookie)
}

func GetAuthCookie(c echo.Context) (*http.Cookie, error) {
	return c.Cookie(AUTH_COOKIE_NAME)
}

func RedirectIfAuthorized(c echo.Context, redirectTo string) error {
	record, ok := GetAuthedUserRecord(c)
	if ok && record != nil {
		if redirectTo != "" {
			return c.Redirect(http.StatusSeeOther, redirectTo)
		}

		return c.Redirect(http.StatusSeeOther, HOME_ROUTE)
	}

	return nil
}

func GetCookie(c echo.Context) (string, error) {
	authCookie, err := c.Cookie(AUTH_COOKIE_NAME)
	if err != nil {
		return "", err
	}

	if authCookie == nil {
		return "", nil
	}

	return authCookie.Value, nil
}

func SetCookie(c echo.Context, value string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = AUTH_COOKIE_NAME
	cookie.Value = value
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return cookie
}
