package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

var cfg config

func main() {
	flag.StringVar(&cfg.addr, "addr", "localhost:7777", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse() // parse input flags, if not will stay as default vals

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{errorLog: errorLog, infoLog: infoLog}

	mux := app.routes()
	infoLog.Printf("starting server on %s", cfg.addr)
	svr := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	err := svr.ListenAndServe()
	errorLog.Fatal(err)
}
