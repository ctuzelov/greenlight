package main

import (
	"fmt"
	"net/http"
)

func (app *application) listMoviesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "In process...")
}

func (app *application) editMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "In process...")
}

func (app *application) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "In process...")
}
