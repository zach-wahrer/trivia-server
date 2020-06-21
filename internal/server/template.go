package server

import "html/template"

const temp1 = `
<!doctype html>

<html lang="en">

    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">

        <title>Trivia!</title>
        <meta name="description" content="Trivia!">
        <meta name="author" content="Zach Wahrer">
    </head>

    <body>
        <div class="container text-center">
            <div class="row pt-4">
                <div class="col">
                        Category: {{.Category}}
                        <br>
                        Difficulty: {{.Difficulty}}
                </div>
            </div>
            <div class="row pt-4">
                <div class="col">
                    <p class="h4">
                        Question: {{.Question}}
                    </p>
                </div>
            </div>
            <div class="row pt-4">
                <div class="col">
                    <div id="answer" style="display:none;">
                        <p class="h4">
                            <em>{{.CorrectAnswer}}</em>
                        <p>
                    </div>
                </div>
            </div>
            <div class="row pt-4">
                <div class="col">
                    <button type="button" class="btn btn-danger" onclick="toggleAnswer()">Show Answer</button>
                </div>
            </div>
            <div class="row pt-4">
                <div class="col">
                    <button type="button" class="btn btn-primary" onclick="window.location.reload();">Next Question</button>
                </div>
            </div>
        </div>

        <script type="text/javascript">
            function toggleAnswer() {
                var id = document.getElementById("answer");
                if (id.style.display === "none") {
                    id.style.display = "block";
                } else {
                    id.style.display = "none";
                }
            }
        </script>

        <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>

    </body>
</html>
    `

var output = template.Must(template.New("trivia").Parse(temp1))
