package main

import (
  "github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/iris"

  "github.com/shubik22/go-robinhood/lib/handlers"
)

func main() {
  iris.Use(recovery.New())
  iris.Get("/users", handlers.UsersHandlerDev)
	iris.Listen(":8080")
}
