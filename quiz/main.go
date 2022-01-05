package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// question struct that stores question with answer
type question struct {
	problem string
	answer  string
}

type quiz struct {
	questions []question
	score     int
}

var (
	qCSV    string
	timeout int
	shuffle bool
)

func init() {
	flag.StringVar(&qCSV, "quiz", "", "A .csv file with questions and answers.")
	flag.IntVar(&timeout, "timeout", 30, "The time limit for answering questions.")
	flag.Bool("shuffle", false, "Select yes to shuffle the quiz")
	flag.Parse()
}

func main() {
	if qCSV == "" {
		log.Fatalln("a .csv file must be provided through -quiz flag")
	}
	questions := readCSV(qCSV)

	qz := quiz{
		questions: questions,
		score:     0,
	}

	if *&shuffle {
		qz.shuffleProblems()
	}

	timeoutCh := time.After((time.Duration(timeout) * time.Second))
	resultCh := make(chan quiz)
	go func() {
		qz.ask()
		resultCh <- qz
	}()
	select {
	case <-resultCh:
	case <-timeoutCh:
		fmt.Println("Timeout!")
	}
	fmt.Println("Score:", qz.score)
}

func (q *quiz) ask() {
	q.score = 0
	for _, question := range q.questions {
		fmt.Println(question.problem)
		var answer string
		fmt.Scanln(&answer)
		if answer == question.answer {
			q.score++
		}
	}
}
func readCSV(s string) []question {
	file, err := os.Open(s)
	if err != nil {
		log.Fatalf("Could not open file : %v \n", err)
	}
	defer file.Close()
	read := csv.NewReader(file)
	records, err := read.ReadAll()
	if err != nil {
		log.Fatalf("Could not parse .csv : %v \n", err)
	}
	var questions []question
	for _, record := range records {
		q := question{problem: record[0], answer: record[1]}
		questions = append(questions, q)
	}
	return questions
}

func (q quiz) shuffleProblems() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(q.questions),
		func(i, j int) {
			q.questions[i], q.questions[j] = q.questions[j], q.questions[i]
		},
	)
}
