// Package trivia interfaces with the opentdb.com API to fetch trivia
package trivia

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"unicode"
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

func GetTrivia() ([]byte, error) {
	resp, err := http.Get(APIURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func ProcessJSON(data []byte) (*Trivia, error) {
	var result Trivia
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}
	result.Data[0].Difficulty = string(unicode.ToUpper(rune(result.Data[0].Difficulty[0]))) + result.Data[0].Difficulty[1:]

	return &result, nil
}

func ProcessQuestion(question string) (result string, err error) {
	return result, err
}
