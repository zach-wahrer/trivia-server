package server

import (
	"fmt"
	"internal/trivia"
	"io"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", IndexHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var body string
	data, err := trivia.GetTrivia()
	if err != nil {
		io.WriteString(w, fmt.Sprintf("Trivia API error: %v. Please try again later.", err))
		return
	}
	body = fmt.Sprintf("%v", data)
	io.WriteString(w, body)
}

func EndPointHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
