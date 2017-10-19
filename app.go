package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type User struct {
	Id   int
	Name string
}

func main() {
	db, _ := gorm.Open("mysql", "root@/test")
	t, _ := template.ParseGlob("views/*")

	router := httprouter.New()
	router.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		t.ExecuteTemplate(w, "index.html", &struct{ Name string }{ps.ByName("name")})
	})
	router.GET("/find/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		result := User{}
		db.First(&result)
		t.ExecuteTemplate(w, "index.html", &struct{ Name string }{result.Name})
	})
	http.ListenAndServe(":8080", router)
}
