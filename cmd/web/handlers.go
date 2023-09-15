package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/heiku-jiqu/snippetapp/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}
	// if use New(), need to be same basename as file
	templ, err := template.ParseFiles(files...)
	if err != nil {
		app.serveError(w, err)
		return
	}

	// retrieve snippets
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serveError(w, err)
		return
	}

	data := templateData{
		Snippets: snippets,
	}

	// "base" is the defined template name, not the filename!
	err = templ.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serveError(w, err)
		return
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serveError(w, err)
		}
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/view.tmpl.html",
	}
	templ, err := template.ParseFiles(files...)
	if err != nil {
		app.serveError(w, err)
		return
	}

	data := &templateData{
		Snippet: snippet,
	}

	err = templ.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serveError(w, err)
		return
	}
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serveError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?&id=%d", id), http.StatusSeeOther)
}
