package handlers

import (
	"github.com/kataras/iris"

	"github.com/shubik22/go-robinhood/lib/client"
)

func AccountsHandler(c *client.Client, ctx *iris.Context) {
	a, _, err := c.Accounts.ListAccounts()
	if err != nil {
		panic(err)
	}
	ctx.JSON(iris.StatusOK, a)
}
