// Package trivia interfaces with the opentdb.com API to fetch trivia
package trivia

import (
	"encoding/json"
	"net/http"
)

const APIURL = "https://opentdb.com/api.php?amount=1"

type Trivia struct {
	ResponseCode int       `json:"response_code"`
	Data         []Results `json:"results"`
}

type Results struct {
	Category      string
	Difficulty    string
	Question      string
	CorrectAnswer string `json:"correct_answer"`
}

func GetTrivia() (*Trivia, error) {
	resp, err := http.Get(APIURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return new(Trivia), nil
}

func ProcessJSON(data string) (*Trivia, error) {
	var result Trivia
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}

	return &result, nil
}
