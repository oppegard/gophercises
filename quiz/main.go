package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	csvPath := flag.String("csv", "problems.csv", "Path to CSV file.")
	flag.Parse()

	records := readRecords(csvPath)
	problems := parseRecords(records)

	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q)

		var input string
		readInput(&input)
		if input == problem.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(records))
}

func readRecords(csvPath *string) [][]string {
	f, err := os.Open(*csvPath)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

type problem struct {
	q string
	a string
}

func parseRecords(records [][]string) []problem {
	ret := make([]problem, len(records))
	for i, record := range records {
		ret[i] = problem{
			q: record[0],
			a: strings.TrimSpace(record[1]),
		}
	}
	return ret
}

func readInput(input *string) {
	n, err := fmt.Scanf("%s\n", input)
	if err != nil {
		log.Fatal(err)
	}
	if n != 1 {
		log.Fatal("Only single-word answers are allowed.")
	}
}
