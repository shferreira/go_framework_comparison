package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"os"
)

func main() {
	facebook := &oauth2.Config{
		ClientID:     os.Getenv("FACEBOOK_KEY"),
		ClientSecret: os.Getenv("FACEBOOK_SECRET"),
		Endpoint: facebook.Endpoint,
		RedirectURL:  "http://localhost:8080/auth/facebook/callback",
		Scopes: []string{
			"email",
		},
	}

	e := echo.New()
	e.Use(middleware.CSRF())
	e.GET("/auth/:provider", func(c echo.Context) error {
		return c.Redirect(302, facebook.AuthCodeURL(c.Get("csrf").(string)))
	})
	e.GET("/auth/:provider/callback", func(c echo.Context) error {
		token, err := facebook.Exchange(oauth2.NoContext, c.QueryParam("code"))
		if err != nil || !token.Valid() || c.QueryParam("state") != c.Cookie("_csrf") {
			return c.String(403, "Invalid token")
		}
		return c.String(200, token.AccessToken)
	})
	e.Start(":8080")
}
