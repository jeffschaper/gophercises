package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"flag"
)

func main() {
	// set optional flags
	var filename string
	var help bool
	
	flag.StringVar(&filename, "csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.BoolVar(&help, "help", false, "usage")
	flag.Parse()
	
	// open the file
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)
	// returns [][]string, which is a slice
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	
	// loop through the file
	var input string
	var correct int
	total := len(records)
	
	for i := 0; i < total; i++ {
		questionSet := records[i]
		// loop through the questions
		for j := 0; j < 1; j++ {
			q := questionSet[j]
			// prompt the user
		 	fmt.Printf("Problem 1: " + q + " = ")
			// wait for a response
			fmt.Scanln(&input)
			// clean input
			// keep track of correct/incorrect answers
			if strings.TrimSuffix(strings.ToLower(input), " ") == questionSet[1] {
				correct++
			}
		}
	}
	// close file
	f.Close()
	// print score
	fmt.Printf("You scored %d out of %d\n", correct, total)
}
