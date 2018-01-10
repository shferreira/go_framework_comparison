package main

import (
  "github.com/gin-gonic/gin"
  "golang.org/x/oauth2"
)

func main() {
  oauthio := &oauth2.Config{
    ClientID:     "oauth_key",
    ClientSecret: "oauth_secret",
    Endpoint: oauth2.Endpoint{
      AuthURL:  "http://localhost:8081/auth",
      TokenURL: "http://localhost:8081/auth/token",
    },
    RedirectURL: "http://localhost:8080/auth/oauthio/callback",
    Scopes: []string{
      "email",
    },
  }

  r := gin.Default()
  r.GET("/auth/oauthio", func(c* gin.Context) {
    c.Redirect(302, oauthio.AuthCodeURL("state"))
  })
  r.GET("/auth/oauthio/callback", func(c* gin.Context) {
    token, err := oauthio.Exchange(oauth2.NoContext, c.Query("code"))
    if err != nil || !token.Valid() {
      c.String(403, "Invalid token"); return
    }
    c.String(200, token.AccessToken)
  })
  r.Run(":8080")
}
