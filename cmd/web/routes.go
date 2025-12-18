package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf)

	mux := http.NewServeMux()

	mux.Handle("GET /{$}", dynamicMiddleware.ThenFunc(app.home))
	mux.Handle("GET /snippet/{id}", dynamicMiddleware.ThenFunc(app.showSnippet))
	mux.Handle("GET /snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippetForm))
	mux.Handle("POST /snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippet))

	mux.Handle("GET  /user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Handle("POST /user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Handle("GET  /user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Handle("POST /user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Handle("POST /user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
