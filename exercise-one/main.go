package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"flag"
)

// problem represents a question and answer pair
type problem struct {
	q string
	a string
}

// setFlags allows the user to enter a custom name of a problem file
func setFlags(filename string) *string {
	// the help flag is included for free
	flag.StringVar(&filename, "csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	return &filename
}

// parseRecords takes in a 2d slice of records and returns a slice of structs
func parseRecords(records [][]string) []problem {
	// make new slice of type slice with length of records
	ret := make([]problem, len(records))
	// range iterates over a slice or a map
	for i, record := range records {
		ret[i] = problem {
			q: record[0],
			a: record[1],
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
		if strings.TrimSpace(strings.ToLower(input)) == problem.a {
			correct++
		}
	}
	return correct
}

func main() {
	// set optional flags
	var filename string
	file := setFlags(filename)
	
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

	score := quiz(problems)

	// print score
	fmt.Printf("You scored %d out of %d\n", score, len(problems))

}
