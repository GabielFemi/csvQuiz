package quiz

import (
	"encoding/csv"
	gabby_csv "github.com/gabielfemi/csvManip/csv"
	"log"
	"os"
)

func ReadQuizFromCsv() {
	csvFile := "quiz.csv"
	gabby_csv.ReadCsv(csvFile)
	StartReading(csvFile)
}

func StartReading(csvFile string) {
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

	for _, field := range csvData {

	}
}