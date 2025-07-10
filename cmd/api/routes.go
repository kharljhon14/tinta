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

	router.HandlerFunc(http.MethodGet, "/v1/blogs/:id", app.requiredPermission("blogs:read", app.showBlogHandlder))
	router.HandlerFunc(http.MethodGet, "/v1/blogs/", app.requiredPermission("blogs:read", app.listBlogsHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/blogs/:id", app.requiredPermission("blogs:write", app.updateBlogHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/blogs/:id", app.requiredPermission("blogs:write", app.deleteBlogHandler))
	router.HandlerFunc(http.MethodPost, "/v1/blogs", app.requiredPermission("blogs:write", app.createBlogHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activate", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.autenticate(router))))
}
