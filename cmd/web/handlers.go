package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//main controller
func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	app.RenderHTML(w, "home.page.html")

}

// handler for /snippet
func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	//fmt.Fprintf(w, "Display a sepcific snipper (ID %d)....", id)

	snippet, err := app.Database.GetSnippet(id)

	if err != nil {
		app.ServerError(w, err)
	}

	if snippet == nil {
		app.NotFound(w)
		return
	}

	fmt.Fprint(w, snippet)

}

// handler for /snippet/new
func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a new snippet"))
}
