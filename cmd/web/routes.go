package main

import "net/http"

func (app *application) routes() http.Handler {
	// servemux == router
	mux := http.NewServeMux()

	// register directory handler in our servemux (router)
	fileServer := http.FileServer(http.Dir(cfg.staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// register a handler in our servemux (router)
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return secureHeaders(mux)
}
