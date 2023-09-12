package main

import (
	"flag"
	"log"
	"net/http"
)

const (
	defaultServeURL    string = "localhost:7777"
	defaultUiStaticDir string = "./ui/static/"
)

func main() {
	addr := flag.String("addr", defaultServeURL, "HTTP network address")
	flag.Parse() // parse input flags, if not will stay as default vals

	// servemux == router
	mux := http.NewServeMux()
	// register a handler in our servemux (router)
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	fileServer := http.FileServer(http.Dir(defaultUiStaticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
