package server

import (
	"io"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", EndPointHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func EndPointHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}
