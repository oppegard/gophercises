package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvPath := flag.String("csv", "problems.csv", "Path to CSV file.")
	limit := flag.Int("limit", 30, "The time limit for the quiz in seconds.")
	flag.Parse()

	records := readRecords(csvPath)
	problems := parseRecords(records)

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Duration(*limit) * time.Second)
		timeout <- true
	}()

	correct := 0
loop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.q)

		inputCh := make(chan string, 1)
		go readInput(inputCh)
		select {
		case <-timeout:
			break loop
		case input := <-inputCh:
			if input == problem.a {
				correct++
			}
			break
		}
	}
	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(records))
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

func readInput(c chan string) {
	var input string
	n, err := fmt.Scanf("%s\n", &input)
	if err != nil {
		log.Fatal(err)
	}
	if n != 1 {
		log.Fatal("Only single-word answers are allowed.")
	}
	c <- input
}
