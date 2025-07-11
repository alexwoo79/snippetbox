package main

import (
	"net/http"

	"github.com/alexwoo79/snippetbox/ui"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// fileServer := http.FileServer(http.Dir("./ui/static/"))

	// mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))

	mux.Handle("GET /about", dynamic.ThenFunc(app.about))

	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))

	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))

	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))

	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))

	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))

	protected := dynamic.Append(app.requireAuthentication)

	// mux.Handle("GET /{$}", protected.ThenFunc(app.home))
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))

	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))

	mux.Handle("GET /account/view", protected.ThenFunc(app.accountView))

	mux.Handle("GET /account/password/update", protected.ThenFunc(app.accountPasswordUpdate))

	mux.Handle("POST /account/password/update", protected.ThenFunc(app.accountPasswordUpdatePost))

	mux.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
