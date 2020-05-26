package quiz

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func format(slice []Quiz) {
	score := 0
	for quiz := range slice {

		question := slice[quiz].Question
		answer := slice[quiz].Answer
		fmt.Println()
		fmt.Print(question, " ")
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')

		if err == nil {

			newAnswer := strings.TrimPrefix(" ", strings.ToLower(answer))
			newInput := strings.TrimPrefix(" ", strings.ToLower(input))

			if newInput == newAnswer {
				newScore := incrementScore(score)
				fmt.Println(newScore)
			} else {
				fmt.Println("You missed it!")
				fmt.Printf("Your answer was: %s", newInput)
				fmt.Printf("The real answer is: %s\n", newAnswer)

			}
		}
	}

}

func incrementScore(previousScore int) int {
	newScore := previousScore + 1
	return newScore
}