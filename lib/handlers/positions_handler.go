package handlers

import (
	"github.com/kataras/iris"

	"github.com/shubik22/robinhood-client"
)

func PositionsHandler(c *robinhood.Client, ctx *iris.Context) {
	p, _, err := c.Positions.ListPositions()
	if err != nil {
		panic(err)
	}
	ctx.JSON(iris.StatusOK, p)
}
