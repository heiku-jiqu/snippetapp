package main

import (
	"fmt"
	"log"
	"net/http"
)

var serveURL string = "localhost:7777"

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "hello world!")
		return
	} else {
		http.NotFound(w, r)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "snippet view")
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "snippet create")
}

func main() {
	// servemux == router
	mux := http.NewServeMux()
	// register a handler in our servemux (router)
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on " + serveURL)
	err := http.ListenAndServe(serveURL, mux)
	log.Fatal(err)
}
