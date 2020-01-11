package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	csvPath := flag.String("csv", "problems.csv", "Path to CSV file.")
	flag.Parse()

	f, err := os.Open(*csvPath)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	correct := 0
	for i, record := range records {
		question, answer := record[0], record[1]
		fmt.Printf("Problem #%d: %s = ", i+1, question)

		var input string
		n, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatal(err)
		}
		if n != 1 {
			log.Fatal("Only single-word answer are allowed.")
		}

		if input == answer {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(records))
}
