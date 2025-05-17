package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	addr string
	staticDir string
}


func main() {

	var cfg config

	flag.StringVar(&cfg.addr,"addr", ":8080", "HTTP network address")
	flag.StringVar(&cfg.staticDir,"staticDir", "./ui/static/", "Path to static files")
	flag.Parse()

	loggerHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(loggerHandler)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(cfg.staticDir))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /home/{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	
	logger.Info("Starting server on " + cfg.addr)
	err := http.ListenAndServe(cfg.addr, mux)
	if err != nil {
		logger.Error(err.Error())
	}
}
