package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"flag"
	"math/rand"
)

// problem represents a question and answer pair
type problem struct {
	q string
	a string
}

// setFlags allows the user to enter a customize the quiz
func setFlags(filename string, shuffle bool) (*string, *bool) {
	// the help flag is included for free
	flag.StringVar(&filename, "csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.BoolVar(&shuffle, "shuffle", false, "shuffle the quiz order")
	flag.Parse()
	return &filename, &shuffle
}

// shuffleQuiz allows the user to shuffle the order of the questions
func shuffleQuiz(problems []problem) {
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})
}

// parseRecords takes in a 2d slice of records and returns a slice of structs
func parseRecords(records [][]string) []problem {
	// make new slice of type slice with length of records
	ret := make([]problem, len(records))
	// range iterates over a slice or a map
	for i, record := range records {
		ret[i] = problem {
			q: record[0],
			a: strings.TrimSpace(record[1]),
		}
	}
	return ret
}

// quiz takes in a slice of type problem and returns the number of correct answers
func quiz(problems []problem) int {
	correct := 0

	for i, problem := range problems {
		// prompt the user
		fmt.Printf("Problem %d: %s = ", i+1, problem.q)
		var input string
		// wait for a response
		fmt.Scanln(&input)
		// clean input
		// keep track of correct/incorrect answers
		if strings.ToLower(input) == problem.a {
			correct++
		}
	}
	return correct
}

func main() {
	// set optional flags
	var filename string
	var shuffle bool
	file, tf := setFlags(filename, shuffle)

	// open the file
	f, err := os.Open(*file)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
		os.Exit(1)
	}

	// read the file
	r := csv.NewReader(f)
	// returns [][]string, which is a 2d slice
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
		os.Exit(1)
	}

	// problems is now a slice of structs
	problems := parseRecords(records)

	if *tf != false && len(problems) > 0 {
		shuffleQuiz(problems)
	}

	score := quiz(problems)

	// print score
	fmt.Printf("You scored %d out of %d\n", score, len(problems))

}
