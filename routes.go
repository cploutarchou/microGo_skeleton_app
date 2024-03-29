package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

//routes The app routes.
func (a *app) routes() *chi.Mux {
	// Add routes here
	a.get("/", a.Handlers.Home)

	// Routes for static files.
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
