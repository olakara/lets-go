package handlers

import (
	"net/http"
)

func (app *Application) snippetCreateHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Info(r.URL.Path)
	addCommonHeaders(w)
	_, err := w.Write([]byte("Create a new snippet"))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *Application) snippetCreatePostHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Info(r.URL.Path)
	addCommonHeaders(w)
	w.WriteHeader(http.StatusCreated)
	// Simulate creating a new snippet
	_, err := w.Write([]byte("Snippet created"))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
