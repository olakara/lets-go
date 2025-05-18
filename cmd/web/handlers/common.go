package handlers

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

type Application struct {
	Logger *slog.Logger
}

func addCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Server", "Go")
}

func (app *Application) serverError(w http.ResponseWriter, r *http.Request, err error) {

	var (
		method = r.Method
		uri    = r.URL.Path
		trace  = string(debug.Stack())
	)

	app.Logger.Error(err.Error(), "method", method, "url", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
