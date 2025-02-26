package main

import (
	"fmt"
	"strings"
)

type Question struct {
	Question string
	Options  []string
	Answer   string
}

var questions = []Question{
	{
		Question: "What is the capital of France?",
		Options:  []string{"Berlin", "Madrid", "Paris", "Rome"},
		Answer:   "Paris",
	},
	{
		Question: "Which language is Go written in?",
		Options:  []string{"Python", "Go", "Java", "C++"},
		Answer:   "Go",
	},

	{
		Question: "What is the capital of France?",
		Options:  []string{"Berlin", "Madrid", "Paris", "Rome"},
		Answer:   "Paris",
	},
}

func main() {
	var score int

	fmt.Println("Welcome to the Quiz! Answer the following questions:")

	// Iterate through each question
	for i, q := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Question)
		fmt.Println("Options:")
		for j, option := range q.Options {
			fmt.Printf("%d. %s\n", j+1, option)
		}

		// Get the user's answer
		var answer string
		fmt.Print("Your answer: ")
		fmt.Scanln(&answer)

		// Check if the answer is correct
		if strings.TrimSpace(strings.ToLower(answer)) == strings.ToLower(q.Answer) {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect. The correct answer was:", q.Answer)
		}
	}

	// Display final score
	fmt.Printf("\nQuiz Finished! Your score is: %d/%d\n", score, len(questions))
}
