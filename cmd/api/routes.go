package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/blogs/:id", app.showBlogHandlder)
	router.HandlerFunc(http.MethodPost, "/v1/blogs", app.createBlogHandler)

	return router
}
