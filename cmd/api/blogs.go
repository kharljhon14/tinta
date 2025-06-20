package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kharljhon14/tinta/internal/data"
)

func (app *application) createBlogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create Blog")
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
