package quiz

import (
	"encoding/csv"
	gabby_csv "github.com/gabielfemi/csvManip/csv"
	"log"
	"os"
	"time"
)

func StartQuiz(csvFile string) {
	log.Println("Started quiz @ ", time.Now())
	format(ReadQuizFromCsv(csvFile))

}
func ReadQuizFromCsv(csvFile string) []Quiz {
	gabby_csv.ReadCsv(csvFile)
	quizzes := startReading(csvFile)
	return quizzes
}

func startReading(csvFile string) []Quiz {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatalln(err)
	}
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	reader.Comma = ';'

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	var quiz Quiz
	var quizzes []Quiz
	for _, field := range csvData {
		quiz.Question = field[0]
		quiz.Answer = field[1]

		// append it into quizzes slice
		quizzes = append(quizzes, quiz)
	}

	return quizzes

}
