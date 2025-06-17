package main

import (
	"fmt"
	"net/http"
)

func (app *application) createBlogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create Blog")
}

func (app *application) showBlogHandlder(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the blog with id: %d\n", id)
}
