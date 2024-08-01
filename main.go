package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type question struct {
	content string
	result  string
}

type quiz struct {
	questions []question
	score     int
}

func (q *quiz) query() {
	for index, question := range q.questions {
		fmt.Printf("Question number %v:\n", (index + 1))
		fmt.Println(question.content, "?")
		var answer string
		fmt.Scanln(&answer)
		if answer == question.result {
			q.score++
		}
	}
}

func main() {
	// Get the appropriate problems CSV file
	fmt.Println("Welcome! Would you like to use the default path for the CSV file? (y/n)")

	var finalPath string

	for {
		var defaultPath string
		fmt.Scanln(&defaultPath)
		if defaultPath == "y" || defaultPath == "n" {
			finalPath = findPathCSV(defaultPath)
			break
		}
		fmt.Println("Unrecognized argument, try again:")
	}

	// Read the CSV file and create the questions slice

	questions := readCSV(finalPath)
	quiz := quiz{questions, 0}

	// Ask questions. This is done with a method from a struct, as to be able to modify the quiz object with its according pointer
	fmt.Println("The quiz will now begin!")
	quiz.query()

	// Reveal the results
	fmt.Printf("From a total of %v questions, you have answered %v correctly.", len(quiz.questions), quiz.score)

}

func findPathCSV(defaultPath string) string {
	//Default file
	if defaultPath == "y" {
		return filepath.Join("./", "problems.csv")
	}
	//Non-Default path
	fmt.Println("Please enter a path for the file:")
	for {
		var customPath string
		fmt.Scanln(&customPath)
		return customPath
	}
}

func readCSV(filePath string) []question {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error while reading the file", err)
		os.Exit(1)
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
		os.Exit(1)
	}

	var questions []question

	for _, record := range records {
		q := question{record[0], record[1]}
		questions = append(questions, q)
	}

	return questions
}
