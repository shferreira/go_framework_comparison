package main

import (
	"github.com/labstack/echo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/instagram"
	"github.com/markbates/goth/providers/linkedin"
	"github.com/markbates/goth/providers/twitter"
	"html/template"
	"os"
)

func main() {
	goth.UseProviders(
		facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), "http://localhost:8080/auth/facebook/callback"),
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:8080/auth/twitter/callback"),
		instagram.New(os.Getenv("INSTAGRAM_KEY"), os.Getenv("INSTAGRAM_SECRET"), "http://localhost:8080/auth/instagram/callback"),
		linkedin.New(os.Getenv("LINKEDIN_KEY"), os.Getenv("LINKEDIN_SECRET"), "http://localhost:8080/auth/linkedin/callback"),
	)

	e := echo.New()
	e.GET("/auth/:provider", func(c echo.Context) error {
		provider, err := goth.GetProvider(c.Param("provider"))
		if err != nil {
			return c.String(404, "Unknown provider")
		}

		session, err := provider.BeginAuth("tmp")
		if err != nil {
			return c.String(500, "Can't begin auth")
		}

		url, err := session.GetAuthURL()
		if err != nil {
			return c.String(500, "Can't get URL")
		}

		return c.Redirect(302, url)
	})
	e.GET("/auth/:provider/callback", func(c echo.Context) error {
		provider, err := goth.GetProvider(c.Param("provider"))
		if err != nil {
			return c.String(404, "Unknown provider")
		}

		session, err := provider.UnmarshalSession("{}")
		if err != nil {
			return c.String(500, "Can't UnmarshalSession")
		}

		_, err = session.Authorize(provider, c.Request().URL.Query())
		if err != nil {
			return c.String(403, "Can't authorize")
		}

		gu, err := provider.FetchUser(session)
		if err != nil {
			return c.String(403, "Can't fetch user")
		}

		t, _ := template.New("foo").Parse(userTemplate)
		return t.ExecuteTemplate(c.Response(), "foo", gu)
	})
	e.GET("/hello/:name", func(c echo.Context) error {
		return c.String(200, "Hello")
	})
	e.Start(":8080")
}

var userTemplate = `
<p><a href="/logout/{{.Provider}}">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`
