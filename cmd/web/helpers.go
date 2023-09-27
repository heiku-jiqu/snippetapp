package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-playground/form/v4"
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

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
		Flash:       app.sessionManager.PopString(r.Context(), "flash"),
	}
}

func (app *application) decodePostForm(r *http.Request, destination interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	err = app.formDecoder.Decode(destination, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError
		// second arg of errors.As needs to be a "double pointer"
		// it points to the pointer that points to form.InvalidDecoderError
		// because *pointer* that points to form.InvalidDecoderError is the one that impls .Error()!
		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
		return err
	}
	return nil
}
