package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "hello world!")
		return
	} else {
		http.NotFound(w, r)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
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

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, fmt.Sprintf("%s Method not allowed", r.Method), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "snippet create")
}
