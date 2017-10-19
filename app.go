package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"html/template"
)

type User struct {
	Id   int
	Name string
}

func main() {
	db, _ := gorm.Open("mysql", "root@/test")
	t := template.Must(template.ParseGlob("views/*.html"))

	e := echo.New()
	e.GET("/hello/:name", func(c echo.Context) error {
		t.ExecuteTemplate(c.Response(), "index.html", &struct{ Name string }{c.Param("name")})
		return nil
	})
	e.GET("/find/:name", func(c echo.Context) error {
		result := User{}
		db.First(&result)
		t.ExecuteTemplate(c.Response(), "index.html", &struct{ Name string }{result.Name})
		return nil
	})
	e.Start(":8080")
}
