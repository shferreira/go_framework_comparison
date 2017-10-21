package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

type User struct {
	Id   int
	Name string
}

func main() {
	db, _ := gorm.Open("mysql", "root@/test")

	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/hello/{name:string}", func(c iris.Context) {
		c.ViewData("Name", c.Params().Get("name"))
		c.View("index.html")
	})
	app.Get("/find/{name:string}", func(c iris.Context) {
		result := User{}
		db.First(&result)
		c.ViewData("Name", result.Name)
		c.View("index.html")
	})
	app.Run(iris.Addr(":8080"))
}
