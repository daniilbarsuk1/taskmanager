package controllers

import (
	"encoding/json"

	"github.com/revel/revel"
)

type Security struct {
	*revel.Controller
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c Security) Signup() revel.Result {
	var user User
	err := json.Unmarshal(c.Params.JSON, &user)
	if err != nil {
		panic(err)
	}

	return c.Redirect("/login")
}
