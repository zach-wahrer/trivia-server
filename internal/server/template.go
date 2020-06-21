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
    <strong>Show answer</strong>
</body>
</html>
    `

var output = template.Must(template.New("trivia").Parse(temp1))
