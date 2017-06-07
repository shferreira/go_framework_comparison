package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id   int
	Name string
}

func main() {
	db, _ := gorm.Open("mysql", "root@/test")

	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/hello/:name", func(c *gin.Context) {
		c.HTML(200, "index.html", &struct{ Name string }{c.Param("name")})
	})
	r.GET("/find/:name", func(c *gin.Context) {
		result := User{}
		db.First(&result)
		c.HTML(200, "index.html", &struct{ Name string }{result.Name})
	})
	r.Run()
}
