package quiz

import (
	"bufio"
	"fmt"
	"os"
)

func format(slice []Quiz) {
	for quiz := range slice {

		question := slice[quiz].Question
		answer := slice[quiz].Answer
		fmt.Println()
		fmt.Print(question, " ")
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')

		if err == nil {
			fmt.Printf("Your answer was: %s", input)
			fmt.Printf("The real answer is: %s\n", answer)
		}
	}

}

func formatAnswer() {

}

func storeInHashMap() {

}
