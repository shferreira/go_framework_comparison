package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(&struct{ Name string }{ "Test" })
	})
	app.Run(iris.Addr("localhost:3001"))
}
