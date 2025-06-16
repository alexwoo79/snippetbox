package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)
	// use the http.NewSeveMux() function to initialize a new servemux ,then
	// register the home function as the handler for "/"url pattern
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))

	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))

	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.snippetCreate))

	// print a log message to say that the server is starting
	// create the new route ,which is restrict to Post Request only
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	// using alice to chain the middleware

	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))

	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))

	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))

	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))

	mux.Handle("POST /user/logout", dynamic.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
