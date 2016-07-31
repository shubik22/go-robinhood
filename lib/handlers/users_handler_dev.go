package handlers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kataras/iris"

	"github.com/shubik22/robinhood-client"
)

type UsersResponse struct {
	Users []robinhood.User
}

func UsersHandlerDev(c *iris.Context) {
	b, err := ioutil.ReadFile("./fixtures/users.json")

	if err != nil {
		panic(err)
	}

	var ur UsersResponse

	if err := json.Unmarshal(b, &ur); err != nil {
		panic(err)
	}

	c.JSON(iris.StatusOK, ur)
}
