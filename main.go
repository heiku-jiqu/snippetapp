package main

import (
	"fmt"
	"log"
	"net/http"
)

var serveURL string = "localhost:7777"

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!")
}

func main() {
	// servemux == router
	mux := http.NewServeMux()
	// register a handler in our servemux (router)
	mux.HandleFunc("/", home)

	log.Print("starting server on " + serveURL)
	err := http.ListenAndServe(serveURL, mux)
	log.Fatal(err)
}
