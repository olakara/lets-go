package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func addCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Server", "Go")
}

func home(w http.ResponseWriter, r *http.Request) {
	addCommonHeaders(w)
	log.Println(r.URL.Path)

	files:= []string{
		"ui/html/base.tmpl",
		"ui/html/partials/nav.tmpl",
		"ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w,"base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}	
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	addCommonHeaders(w)
	log.Println(r.URL.Path)
	// Extract the snippet ID from the URL
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := "Display a specific snippet with ID: " + strconv.Itoa(id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	addCommonHeaders(w)
	w.Write([]byte("Create a new snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	addCommonHeaders(w)
	w.WriteHeader(http.StatusCreated)
	// Simulate creating a new snippet
	w.Write([]byte("Snippet created"))
}
