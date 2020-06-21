package server

import "html/template"

const temp1 = `
<!doctype html>

<html lang="en">
<head>
    <meta charset="utf-8">

    <title>Trivia!</title>
    <meta name="description" content="Trivia!">
    <meta name="author" content="Zach Wahrer">
</head>

<body>
    <strong>Category:</strong> {{.Category}}<br>
    <strong>Difficulty:</strong> {{.Difficulty}}
    <br><br>
    <strong>Question:</strong> {{.Question}}

    <br><br>


    <div id="answer" style="display:none;" padding=10>
        <h3>{{.CorrectAnswer}}</h3>
    </div>
    <br><br>

    <button onclick="toggleAnswer()">Show Answer</button>

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
</body>
</html>
    `

var output = template.Must(template.New("trivia").Parse(temp1))
