package main

import (
	"log"
	"net/http"
)

var serveURL string = "localhost:7777"

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
