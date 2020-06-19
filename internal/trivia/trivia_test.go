package trivia

import (
	"testing"
)

func TestAPICall(t *testing.T) {
	trivia, err := GetTrivia()
	if err != nil {
		t.Fatal(err)
	}

	if trivia.ResponseCode != 0 {
		t.Errorf("triviaAPI returned unexpected result: got code %d", trivia.ResponseCode)
	}

	if len(trivia.Results) == 0 {
		t.Errorf("triviaAPI returned empty result")
	}

}
