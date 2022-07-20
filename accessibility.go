package main

import (
	"net/http"
)

//get helper function for app get request.
func (a *app) get(s string, h http.HandlerFunc) {
	a.App.Routes.Get(s, h)
}

//get helper function for app post request.
func (a *app) post(s string, h http.HandlerFunc) {
	a.App.Routes.Post(s, h)

}

//get helper function for middlewares
func (a *app) use(m ...func(handler http.Handler) http.Handler) {
	a.App.Routes.Use(m...)
}
