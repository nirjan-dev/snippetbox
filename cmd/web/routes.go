package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := http.NewServeMux()

	mux.Handle("GET /{$}", dynamicMiddleware.Then(http.HandlerFunc(app.home)))
	mux.Handle("GET /snippet/{id}", dynamicMiddleware.Then(http.HandlerFunc(app.showSnippet)))
	mux.Handle("POST /snippet/create", dynamicMiddleware.Then(http.HandlerFunc(app.createSnippet)))
	mux.Handle("GET /snippet/create", dynamicMiddleware.Then(http.HandlerFunc(app.createSnippetForm)))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
