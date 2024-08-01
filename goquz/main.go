package main

import (
	"github.com/babugeet/goquz/goquz/internal/constants"
	"github.com/babugeet/goquz/goquz/internal/quizutils"
)

func main() {

	csvContent, _ := quizutils.ReadCSVfile(constants.Filename)
	quizutils.QuizQA(csvContent)

}
