//go:generate qtc -dir=views
package main

import (
	"./views"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type User struct {
	Id   int
	Name string
}

func main() {
	db, _ := gorm.Open("mysql", "root@/test")

	e := echo.New()
	e.GET("/hello/:name", func(c echo.Context) error {
		views.WritePageTemplate(c.Response(), &views.Index{c.Param("name")})
		return nil
	})
	e.GET("/find/:name", func(c echo.Context) error {
		result := User{}
		db.First(&result)
		views.WritePageTemplate(c.Response(), &views.Index{result.Name})
		return nil
	})
	e.Start(":8080")
}
