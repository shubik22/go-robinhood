package main

import (
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/iris"

	"github.com/shubik22/go-robinhood/lib/client"
	"github.com/shubik22/go-robinhood/lib/handlers"
)

type page struct {
	Title string
}

func main() {
	c := client.NewClient()

	iris.Use(recovery.New())
	iris.Get("/users", handlers.UsersHandlerDev)
	iris.Get("/accounts", func(ctx *iris.Context) {
		handlers.AccountsHandler(c, ctx)
	})

	iris.Static("/static_files/css", "./static_files/css", 1)
	iris.Static("/static_files/js", "./static_files/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.MustRender("/static_files/html/index.html", page{Title: "Home"})
	})

	iris.Listen(":8080")
}
