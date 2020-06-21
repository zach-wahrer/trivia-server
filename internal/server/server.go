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
	resp, err := trivia.GetTrivia()
	if err != nil {
		PrintError(w, fmt.Errorf("TriviaAPI: %s", err))
		return
	}
	data, err := trivia.ProcessJSON(resp)
	if err != nil {
		PrintError(w, fmt.Errorf("ProcessJSON: %s", err))
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

func PrintError(w http.ResponseWriter, err error) {
	io.WriteString(w, fmt.Sprintf("Error: %v. Please try again later.", err))
}
