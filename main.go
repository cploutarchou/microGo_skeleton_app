package main

import (
	"app/data"
	"app/handlers"
	"app/middleware"
	microGo "github.com/cploutarchou/MicroGO"
)

type app struct {
	App        *microGo.MicroGo
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
