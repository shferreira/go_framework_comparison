package main

import (
	"github.com/buaazp/fasthttprouter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/valyala/fasthttp"
	"html/template"
)

type User struct {
	Id   int
	Name string
}

func main() {
	db, _ := gorm.Open("mysql", "root@/test")
	t := template.Must(template.ParseGlob("views/*.html"))

	router := fasthttprouter.New()
	router.GET("/hello/:name", func(c *fasthttp.RequestCtx) {
		c.SetContentType("text/html")
    t.ExecuteTemplate(c, "index.html", &struct{ Name string }{c.UserValue("name").(string)})
	})
	router.GET("/find/:name", func(c *fasthttp.RequestCtx) {
		result := User{}
		db.First(&result)
    c.SetContentType("text/html")
		t.ExecuteTemplate(c, "index.html", &struct{ Name string }{result.Name})
	})
	fasthttp.ListenAndServe(":8080", router.Handler)
}
