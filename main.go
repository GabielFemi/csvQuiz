package main

import (
	"fmt"
	"github.com/gabielfemi/csvQuiz/quiz"
)

func main() {
	quizFromCsv := quiz.ReadQuizFromCsv("quiz.csv")
	fmt.Println(quizFromCsv)
}