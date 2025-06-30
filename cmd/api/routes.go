package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/blogs/:id", app.showBlogHandlder)
	router.HandlerFunc(http.MethodGet, "/v1/blogs/", app.listBlogsHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/blogs/:id", app.updateBlogHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/blogs/:id", app.deleteBlogHandler)
	router.HandlerFunc(http.MethodPost, "/v1/blogs", app.createBlogHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)

	return app.rateLimit(app.recoverPanic(router))
}
