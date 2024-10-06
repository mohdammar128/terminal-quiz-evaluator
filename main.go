package main

import (
	"calculator/quiz"
)

func main() {
	filePath := "quiz.csv"
	quiz.ReadCsv(filePath)
	q := quiz.NewQuiz()

}
