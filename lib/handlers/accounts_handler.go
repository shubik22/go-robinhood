package handlers

import (
	"github.com/kataras/iris"

	"github.com/shubik22/robinhood-client"
)

func AccountsHandler(c *robinhood.Client, ctx *iris.Context) {
	a, _, err := c.Accounts.ListAccounts()
	if err != nil {
		panic(err)
	}
	ctx.JSON(iris.StatusOK, a)
}
