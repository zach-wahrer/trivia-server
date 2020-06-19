module github.com/zachtheclimber/trivia-server

go 1.13

require internal/server v1.0.0

replace internal/server => ./internal/server

require internal/trivia v1.0.0

replace internal/trivia => ./internal/trivia
