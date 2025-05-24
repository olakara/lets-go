package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"snippetbox/cmd/web/handlers"

	"github.com/jackc/pgx/v5/pgxpool"
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

	dbPool, err:= pgxpool.New(context.Background(), "postgres://spinUser:jaCk-will-c0st@localhost:5432/snippetDB")
	if err != nil {
		logger.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}	
	defer dbPool.Close()

	app := handlers.NewApplication(logger, dbPool)

	logger.Info("Starting server on " + cfg.addr)
	err = http.ListenAndServe(cfg.addr, app.Routes())
	if err != nil {
		logger.Error(err.Error())
	}
}
