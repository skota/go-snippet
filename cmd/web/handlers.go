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
	//w.Write([]byte("Display a specific snippet"))
	fmt.Fprintf(w, "Display a sepcific snipper (ID %d)....", id)
}

// handler for /snippet/new
func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a new snippet"))
}
