package handlers

import (
	"html/template"
	"net/http"
)

func (app *Application) homeHandler(w http.ResponseWriter, r *http.Request) {

	app.Logger.Info(r.URL.Path)
	addCommonHeaders(w)

	files := []string{
		"ui/html/base.tmpl",
		"ui/html/partials/nav.tmpl",
		"ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
