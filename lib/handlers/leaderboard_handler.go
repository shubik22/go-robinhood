package handlers

import (
  "github.com/kataras/iris"

  "github.com/shubik22/go-robinhood/lib/app"
)

func LeaderboardHandler(a *app.App, ctx *iris.Context) {
  lb := a.GetLeaderboard()
  ctx.JSON(iris.StatusOK, lb)
}
