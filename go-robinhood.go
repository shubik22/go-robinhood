package main

import (
	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/iris"

	"github.com/shubik22/go-robinhood/lib/client"
	"github.com/shubik22/go-robinhood/lib/handlers"
)

func main() {
	c := client.NewClient()

	iris.Use(logger.New(iris.Logger))
	iris.Use(recovery.New())

	iris.Get("/users", handlers.UsersHandlerDev)
	iris.Get("/accounts", func(ctx *iris.Context) {
		handlers.AccountsHandler(c, ctx)
	})

	iris.Static("/css", "./static_files/css", 1)
	iris.Static("/js", "./static_files/js", 1)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.ServeFile("./static_files/html/index.html", false)
	})

	iris.Listen(":8080")
}
