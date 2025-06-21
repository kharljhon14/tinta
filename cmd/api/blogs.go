package main

import (
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

	v := validator.New()

	v.Check(input.Title != "", "title", "must be provided")
	v.Check(len(input.Title) <= 255, "title", "must not be more than 255 bytes long")

	v.Check(input.Content != "", "content", "must be provided")
	v.Check(len(input.Content) <= 500, "content", "must not be more than 500 bytes long")

	v.Check(input.Tags != nil, "tags", "must be provided")
	v.Check(len(input.Tags) >= 0, "tags", "must contain at least 1 tag")
	v.Check(len(input.Tags) <= 5, "tags", "must not contain more than 5 genres")
	v.Check(validator.Unique(input.Tags), "tags", "must not contain duplicate values")

	if !v.Valid() {
		app.faildValidationResponse(w, r, v.Errors)
		return
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showBlogHandlder(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	blog := data.Blog{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Sample Title",
		Author:    "Kharl",
		Tags:      []string{"Tech", "Guide", "Web"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"blog": blog}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
