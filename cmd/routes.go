package main

import (
	"net/http"
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

// These routes are accessible with https://ip_address through ingress/letsencrypt or with http//localhost in PC 
func (app *application) webRoutes() http.Handler {

    standardMiddleware := alice.New(/*app.recoverPanic, app.logRequest, secureHeaders */)

	// Initialize a new router, then register the functions as the handler for the URL patterns.
	// Pat matches patterns in the order that they are registered.
	mux := pat.New()
    mux.Get("/ping", http.HandlerFunc(app.ping))

	// Create a file server which serves files out of the "./ui/static" directory.
	// Path given to the http.Dir function is relative to the project directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register the file server as the handler for all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}