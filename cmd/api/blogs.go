package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/kharljhon14/tinta/internal/data"
	"github.com/kharljhon14/tinta/internal/validator"
)

func (app *application) createBlogHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title   string   `json:"title"`
		Content string   `json:"content"`
		Tags    []string `json:"tags"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	// TODO: Get author from request
	blog := &data.Blog{

		Content:   input.Content,
		Title:     input.Title,
		Tags:      input.Tags,
		Author:    "Kharl",
		Version:   1,
		CreatedAt: time.Now(),
	}

	if data.ValidateBlog(v, blog); !v.Valid() {
		app.faildValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Blogs.Insert(blog)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/blogs/%d", blog.ID))

	err = app.writeJSON(w, http.StatusOK, envelope{"blog": blog}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showBlogHandlder(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	blog, err := app.models.Blogs.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErroRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"blog": blog}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
