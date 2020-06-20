// Package trivia interfaces with the opentdb.com API to fetch trivia
package trivia

import "net/http"

const APIURL = "https://opentdb.com/api.php?amount=1"

type Trivia struct {
	ResponseCode int
	Results      map[string]string
}

func GetTrivia() (*Trivia, error) {
	resp, err := http.Get(APIURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return new(Trivia), nil
}

func ProcessJSON(data []byte) (*Trivia, error) {

	return new(Trivia), nil
}
