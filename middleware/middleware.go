package middleware

import (
	"app/data"
	microGo "github.com/cploutarchou/MicroGO"
)

type Middleware struct {
	App    *microGo.MicroGo
	Models data.Models
}
