package main

import (
	"encoding/csv"
	"fmt"
	"io"
	_"io"
	"log"
	"os"
	"time"
)

func errorHandler(err error){
	if err != nil {
		log.Fatal(err)
	}
	return
}

func timer(){
	var starter string
	fmt.Println("Press any key to start")
	_, err := fmt.Scan(&starter)
	errorHandler(err)
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

func makeScore(arr []string, questions int){
	answerCount := 0
	for _ = range arr {
		answerCount += 1
	}
	fmt.Println("Passed ", answerCount, " out of ", questions)
}

func main(){
	var correct []string
	f, err := os.OpenFile("problems.csv", os.O_RDWR|os.O_CREATE, 0755) // Open file
	errorHandler(err)
	read := csv.NewReader(f)
	questionCount := 0
	for {
		record, err := read.Read()
		if err == io.EOF {
			break
		}
		if answer := scanAnswers(record[0], record[1]); answer != "" {
			correct = append(correct, answer)
		}
		questionCount += 1
		errorHandler(err)
	}
	f.Close()
	makeScore(correct, questionCount)
}