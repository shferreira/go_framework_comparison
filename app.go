package main

import (
	"github.com/labstack/echo"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	users := map[int]*User{}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, users)
	})
	e.PUT("/:id", func(c echo.Context) error {
		user := &User{}
		c.Bind(user)
		users[user.Id] = user
		return c.JSON(200, user)
	})
	e.Start(":8080")
}
