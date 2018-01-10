package main

import (
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/auth", func(c* gin.Context) {
    c.Redirect(302, c.Query("redirect_uri") + "?code=12345&state=" + c.Query("state"));
  });
  r.POST("/auth/token", func(c* gin.Context) {
    c.JSON(200, gin.H{ "access_token": "54321" });
  });
  r.Run(":8081")
}
