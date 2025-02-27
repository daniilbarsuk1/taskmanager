package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Signup() revel.Result {
	return c.Render()
}

func (c App) Home() revel.Result {
	return c.Render()
}

func (c App) Login() revel.Result {
	return c.Render()
}
