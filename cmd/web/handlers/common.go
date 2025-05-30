package handlers

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

type Application struct {
	logger *slog.Logger
}

func NewApplication(logger *slog.Logger) *Application {
	return &Application{
		logger: logger,
	}
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

	app.logger.Error(err.Error(), "method", method, "url", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/home/", app.homeHandler)
	mux.HandleFunc("/snippet/view/{id}", app.snippetViewHandler)
	mux.HandleFunc("GET /snippet/create", app.snippetCreateHandler)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePostHandler)

	return mux
}
