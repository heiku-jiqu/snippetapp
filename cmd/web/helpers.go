package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serveError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, status int, templateName string, data *templateData) {
	templ, ok := app.templateCache[templateName]
	if !ok {
		app.serveError(w, fmt.Errorf("Failed to get template %s", templateName))
		return
	}

	buf := new(bytes.Buffer)
	err := templ.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serveError(w, err)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}
