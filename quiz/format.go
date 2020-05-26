package quiz

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func format(slice []Quiz) {
	score := 0
	fmt.Println("Welcome to csvQuiz written by Akinyosoye Gabriel. :)")
	fmt.Println()

	for quiz := range slice {

		question := slice[quiz].Question
		answer := slice[quiz].Answer
		fmt.Print(question, " ")
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')

		if err == nil {
			switch input {
			case answer+"\r\n": fmt.Println("Correct!"); score += 1; fmt.Println()
			default : fmt.Println("Incorrect! Answer is", answer);fmt.Println()

			}

		} else {
			log.Fatalln(err)
		}
	}
	fmt.Println()
	fmt.Printf("Your total score is %d/%d", score, len(slice))

}

func incrementScore(previousScore int) int {
	newScore := previousScore + 1
	return newScore
}