package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
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

	fmt.Println("Would you like to use the default timer? (y/n)")
	timerSeconds := 30 //Default
	for {
		var timerDefault string
		fmt.Scanln(&timerDefault)
		if timerDefault == "y" {
			break
		} else if timerDefault == "n" {
			fmt.Println("How many seconds would you like to have in total?")
			var customTimer string
			fmt.Scanln(&customTimer)
			timerSeconds, _ = strconv.Atoi(customTimer)
			break
		}
		fmt.Println("Unrecognized argument, try again:")
	}

	// Read the CSV file and create the questions slice

	questions := readCSV(finalPath)
	quizInstance := quiz{questions, 0}

	// Ask questions. This is done with a method from a struct, as to be able to modify the quiz object with its according pointer
	fmt.Println("The quiz will begin as soon as you press a button!")
	fmt.Scan()
	fmt.Println(timerSeconds, "seconds, START!")

	// Quiz is quizzed within the bounds of the func of the timeout - https://gobyexample.com/timeouts

	timeoutChannel := make(chan quiz) //accepts an argument of the type quiz, not buffered so it is blocking
	go func() {                       // Will do {}, as long as "cases" dont kick in
		quizInstance.query()
		timeoutChannel <- quizInstance
	}()
	select {
	case <-timeoutChannel: //Finishes normally
	case <-time.After(time.Duration(timerSeconds) * time.Second): //timeout occurs before finishing
		fmt.Println("Time's off!")
	}

	// Reveal the results
	fmt.Printf("From a total of %v questions, you have answered %v correctly.", len(quizInstance.questions), quizInstance.score)

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
