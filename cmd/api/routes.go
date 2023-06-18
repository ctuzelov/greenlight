package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// what about the error messages that httprouter automatically sends when it can’t find
	// a matching route? By default, these will still be the same plain-text (non-JSON) responses
	// that we saw earlier in the book.
	// Fortunately, httprouter allows us to set our own custom error handlers when we initialize
	// the router. These custom handlers must satisfy the http.Handler interface, which is good
	// news for us because it means we can easily re-use the notFoundResponse() and
	// methodNotAllowedResponse() helpers that we just made.
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return router
}
