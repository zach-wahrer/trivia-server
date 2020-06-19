// Package trivia interfaces with the opentdb.com API to fetch trivia
package trivia

const APIURL = "https://opentdb.com/api.php?amount=1"

type Trivia struct {
	ResponseCode int
	Results      map[string]string
}

func GetTrivia() (*Trivia, error) {
	return new(Trivia), nil
}
