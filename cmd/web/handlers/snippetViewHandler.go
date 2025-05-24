package handlers

import (
	"net/http"
	"strconv"
)

func (app *Application) snippetViewHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Info(r.URL.Path)
	addCommonHeaders(w)
	// Extract the snippet ID from the URL
	queryParam := r.PathValue("id")
	id, err := strconv.Atoi(queryParam)
	if err != nil || id < 1 {
		app.serverError(w, r, err)
		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.logger.Info(snippet.ToString())
}
