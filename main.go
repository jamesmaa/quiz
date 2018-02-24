package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var csvFile = flag.String("csv", "problems.csv", "CSV file for problems")

type Problem struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Quiz struct {
	Problems []Problem
	Correct  int
}

func main() {
	// Parse csv
	csvFile, _ := os.Open(*csvFile)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var problems []Problem
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		question := line[0]
		answer := line[1]
		problems = append(problems, Problem{
			Question: question,
			Answer:   answer,
		})
	}

	// Display question
	quiz := Quiz{Problems: problems}
	var response string
	for _, p := range problems {
		fmt.Print(p.Question + ": ")
		fmt.Scanln(&response)
		if response == p.Answer {
			fmt.Println("Correct!")
			quiz.Correct = quiz.Correct + 1
		}
	}

	fmt.Printf("You got %d / %d problems correct\n", quiz.Correct, len(quiz.Problems))
}
