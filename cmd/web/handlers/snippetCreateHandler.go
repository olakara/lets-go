package handlers

import (
	"log"
	"net/http"
)

func SnippetCreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	addCommonHeaders(w)
	w.Write([]byte("Create a new snippet"))
}

func SnippetCreatePostHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	addCommonHeaders(w)
	w.WriteHeader(http.StatusCreated)
	// Simulate creating a new snippet
	w.Write([]byte("Snippet created"))
}
