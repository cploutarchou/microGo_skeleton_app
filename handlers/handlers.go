package handlers

import (
	"app/data"
	microGo "cloud0.christosploutarchou.com/cploutarchou/MicroGO"
	"net/http"
)

type Handlers struct {
	APP    *microGo.MicroGo
	Models data.Models
}

// Home Request Handler for home package.
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.APP.ErrorLog.Println("Unable to render page :", err)
	}
}
