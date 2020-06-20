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

	// if trivia.Question == "" {
	// 	t.Errorf("triviaAPI returned empty result")
	// }

}

func TestProcessJSON(t *testing.T) {
	// data := {"response_code":0,"results":[{"category":"Science: Computers","type":"multiple","difficulty":"medium","question":"What is the name of the default theme that is installed with Windows XP?","correct_answer":"Luna","incorrect_answers":["Neptune","Whistler","Bliss"]}]}`
	data := `{"response_code":5,"test":"this is the test string"}`

	trivia, err := ProcessJSON(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Errorf("%v", trivia)
	// if trivia.Category != "Science" {
	// 	t.Errorf("processJSON did not correctly convert JSON: got \"%s\" want \"Science\"",
	// 		trivia.Category)
	// }

}
