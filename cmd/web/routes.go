package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// use the http.NewSeveMux() function to initialize a new servemux ,then
	// register the home function as the handler for "/"url pattern
	mux.HandleFunc("GET /{$}", app.home)

	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)

	mux.HandleFunc("GET /snippet/create", app.snippetCreate)

	// print a log message to say that the server is starting
	// create the new route ,which is restrict to Post Request only
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// using alice to chain the middleware
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
