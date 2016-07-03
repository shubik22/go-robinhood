package handlers

import (
  "encoding/json"
  "io/ioutil"

  "github.com/kataras/iris"

  "github.com/shubik22/go-robinhood/lib/models"
)

type UsersResponse struct {
  Users []models.User
}

func UsersHandler(c *iris.Context) {
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
