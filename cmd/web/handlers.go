package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
		app.errorLog.Println(err.Error())
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		return
	}

	// "base" is the defined template name, not the filename!
	err = templ.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "failed to execute template", http.StatusInternalServerError)
		return
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.NotFound(w, r)
		fmt.Fprintf(w, "failed to parse id: %q", idString)
		return
	}
	if id < 1 {
		fmt.Fprintf(w, "expected id greater than 0, but got %d", id)
		return
	}

	fmt.Fprintf(w, "snippet view id: %d", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, fmt.Sprintf("%s Method not allowed", r.Method), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "snippet create")
}
