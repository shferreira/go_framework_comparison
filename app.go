//go:generate ego views
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"./views"
)

type User struct {
	Id   int
	Name string
}

func main() {
	db, _ := gorm.Open("mysql", "root@/test")

	e := echo.New()
	e.GET("/hello/:name", func(c echo.Context) error {
		return views.Index(c.Response(), c.Param("name"))
	})
	e.GET("/find/:name", func(c echo.Context) error {
		result := User{}
		db.First(&result)
		return views.Index(c.Response(), result.Name)
	})
	e.Start(":8080")
}
