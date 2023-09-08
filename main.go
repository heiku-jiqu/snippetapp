package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!")
}

func main() {
	// servemux == router
	mux := http.NewServeMux()
	// register a handler in our servemux (router)
	mux.HandleFunc("/", home)

	log.Print("starting server on :7777")
	err := http.ListenAndServe(":7777", mux)
	log.Fatal(err)
}
