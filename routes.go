package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

//routes The app routes.
func (a *app) routes() *chi.Mux {
	// Add routes here
	a.get("/", a.Handlers.Home)
	a.get("/sessions", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/users/login", func(w http.ResponseWriter, r *http.Request) {})
	a.post("/users/login", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/users/logout", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/users/forgot-password", func(w http.ResponseWriter, r *http.Request) {})
	a.post("/users/forgot-password", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/users/reset-password", func(w http.ResponseWriter, r *http.Request) {})
	a.post("/users/reset-password", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/form", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/json", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/xml", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/download", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/encryption", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/cache-test", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/cache-test", func(w http.ResponseWriter, r *http.Request) {})
	a.post("/api/save-in-cache", func(w http.ResponseWriter, r *http.Request) {})
	a.post("/api/get-from-cache", func(w http.ResponseWriter, r *http.Request) {})
	a.post("/api/delete-from-cache", func(w http.ResponseWriter, r *http.Request) {})
	a.post("/api/empty-cache", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/test-email", func(w http.ResponseWriter, r *http.Request) {})
	a.App.Routes.Post("/form", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/create_user", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/get_all_users", func(w http.ResponseWriter, r *http.Request) {})

	a.get("/get_user/{id}", func(w http.ResponseWriter, r *http.Request) {})
	a.get("/update_user/{id}", func(w http.ResponseWriter, r *http.Request) {})

	// Routes for static files.
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
