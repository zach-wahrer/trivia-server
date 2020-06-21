// Package trivia interfaces with the opentdb.com API to fetch trivia and process it
package trivia

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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
	result.Data[0].Question = ProcessQuestion(result.Data[0].Question)
	result.Data[0].CorrectAnswer = ProcessAnswer(result.Data[0].CorrectAnswer)
	return &result, nil
}

func ProcessQuestion(question string) (result string) {
	result = FixQuoteText(question)
	endchar := string(result[len(result)-1])
	lastThree := string(result[len(result)-3 : len(result)])
	if endchar != "?" && endchar != ":" && lastThree != "..." {
		result = "(True/False) " + result
	}
	return result
}

func ProcessAnswer(answer string) string {
	return FixQuoteText(answer)
}

func FixQuoteText(text string) (result string) {
	result = strings.Replace(text, "&quot;", "\"", -1)
	result = strings.Replace(result, "&#039;", "'", -1)
	return result
}
