package handlers

import (
	"log"
	"net/http"
)

func SnippetCreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	addCommonHeaders(w)
	_, err := w.Write([]byte("Create a new snippet"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func SnippetCreatePostHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	addCommonHeaders(w)
	w.WriteHeader(http.StatusCreated)
	// Simulate creating a new snippet
	_, err := w.Write([]byte("Snippet created"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
