package main

import (
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	facebook := &oauth2.Config{
		ClientID:     os.Getenv("FACEBOOK_KEY"),
		ClientSecret: os.Getenv("FACEBOOK_SECRET"),
		RedirectURL:  "http://localhost:8080/auth/facebook/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
		Scopes: []string{
			"email",
		},
	}

	e := echo.New()
	e.GET("/auth/:provider", func(c echo.Context) error {
		return c.Redirect(302, facebook.AuthCodeURL("tmp"))
	})
	e.GET("/auth/:provider/callback", func(c echo.Context) error {
		token, err := facebook.Exchange(oauth2.NoContext, c.QueryParam("code"))
		if err != nil || !token.Valid() {
			return c.String(403, "Invalid token")
		}
		res, err := http.Get("https://graph.facebook.com/me?fields=email,first_name,last_name,link,about,id,name,picture,location&access_token=" + token.AccessToken)
		if err != nil {
			return c.String(403, "Can't read user data")
		}
		body, _ := ioutil.ReadAll(res.Body)

		return c.String(200, string(body))
	})
	e.Start(":8080")
}
