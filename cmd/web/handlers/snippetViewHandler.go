package handlers

import (
	"net/http"
	"strconv"
)

func (app *Application) SnippetViewHandler(w http.ResponseWriter, r *http.Request) {
	app.Logger.Info(r.URL.Path)
	addCommonHeaders(w)
	// Extract the snippet ID from the URL
	queryParam := r.PathValue("id")
	id, err := strconv.Atoi(queryParam)
	if err != nil || id < 1 {
		app.serverError(w, r, err)
		return
	}
	msg := "Display a specific snippet with ID: " + strconv.Itoa(id)
	_, err = w.Write([]byte(msg))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
