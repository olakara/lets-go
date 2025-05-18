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

	app := &handlers.Application{
		Logger: logger,
	}

	logger.Info("Starting server on " + cfg.addr)
	err := http.ListenAndServe(cfg.addr, app.Routes())
	if err != nil {
		logger.Error(err.Error())
	}
}
