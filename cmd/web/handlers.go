package main

import (
	"log"
	"net/http"
	"strconv"
)

func addCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Server", "Go")
}

func home(w http.ResponseWriter, r *http.Request) {
	addCommonHeaders(w)
	log.Println(r.URL.Path)
	w.Write([]byte("Hello from Snippetbox"))
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
