package handlers

import (
	"net/http"
)

func addCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Server", "Go")
}
