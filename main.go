package main

import (
	"encoding/csv"
	"fmt"
	"io"
	_ "io"
	"log"
	"os"
	"time"
)

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
	return
}

func scanAnswers(record string, expectedAnswer string) string {
	var answer string

	fmt.Println("What is ", record)
	_, err := fmt.Scan(&answer)
	errorHandler(err)

	if answer == expectedAnswer {
		return answer
	}

	return ""
}

func makeScore(arr []string, questions int) {
	answerCount := 0
	for _ = range arr {
		answerCount += 1
	}
	fmt.Println("\nPassed ", answerCount, " out of ", questions)
}

func readLines() int{
	f, err := os.OpenFile("problems.csv", os.O_RDWR|os.O_CREATE, 0755) // Open file
	errorHandler(err)
	read := csv.NewReader(f)
	csvFile, err := read.ReadAll()
	errorHandler(err)
	return len(csvFile)
}

func main() {
	var correct []string
	var enter string

	fmt.Println("Press s and the enter key to start")
	_, err := fmt.Scan(&enter)
	errorHandler(err)

	timer1 := time.NewTimer(10 * time.Second) // return channel
	defer timer1.Stop() // stop after main func execution
	done := make(chan bool) // create a channel
	go func(chanel chan bool) {
		f, err := os.OpenFile("problems.csv", os.O_RDWR|os.O_CREATE, 0755) // Open file
		errorHandler(err)
		read := csv.NewReader(f)
		for {
			record, err := read.Read()
			if err == io.EOF {
				break
			}
			if answer := scanAnswers(record[0], record[1]); answer != "" {
				correct = append(correct, answer)
			}
			errorHandler(err)
		}
		f.Close()
		done <- true
		makeScore(correct, readLines())
	}(done)
	for {
		select {
		case <-done:
			return
		case <-timer1.C:
			makeScore(correct, readLines())
			return
		}
	}
}
