package handlers

import (
	"log"
	"net/http"
	"strconv"
)

func SnippetViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	addCommonHeaders(w)
	// Extract the snippet ID from the URL
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := "Display a specific snippet with ID: " + strconv.Itoa(id)
	_, err = w.Write([]byte(msg))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
