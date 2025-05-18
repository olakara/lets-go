package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"snippetbox/cmd/web/handlers"
)

type config struct {
	addr      string
	staticDir string
}

func main() {

	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":8080", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "staticDir", "./ui/static/", "Path to static files")
	flag.Parse()

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})

	logger := slog.New(loggerHandler)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(cfg.staticDir))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /home/{$}", handlers.HomeHandler)
	mux.HandleFunc("GET /snippet/view/{id}", handlers.SnippetViewHandler)
	mux.HandleFunc("GET /snippet/create", handlers.SnippetCreateHandler)
	mux.HandleFunc("POST /snippet/create", handlers.SnippetCreatePostHandler)

	logger.Info("Starting server on " + cfg.addr)
	err := http.ListenAndServe(cfg.addr, mux)
	if err != nil {
		logger.Error(err.Error())
	}
}
