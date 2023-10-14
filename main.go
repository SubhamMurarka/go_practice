package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func userInput(question string, done chan bool, ch chan string) {
	fmt.Printf("%v\n", question)
	var answer string
	fmt.Scan(&answer)
	done <- true
	ch <- answer
}

func loadQuestions(address string) (records [][]string, err error) {
	fd, error := os.Open(address)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Successfully opened the csv file")
	defer fd.Close()

	fileReader := csv.NewReader(fd)
	records, error = fileReader.ReadAll()

	return records, err
}

func main() {
	var limit int
	flag.IntVar(&limit, "limit", 4, "change timeout duration")
	var address string
	flag.StringVar(&address, "file", "/home/murarka/go_project/problems.csv", "add different file format (question,answer)")
	flag.Parse()
	records, error := loadQuestions(address)
	if error != nil {
		log.Fatal(error)
	}
	ch := make(chan string)
	var correct int
	var incorrect int
	done := make(chan bool)
	timeout := time.Duration(limit) * time.Second
	fmt.Println("Press the Enter Key to start!")
	fmt.Scanln() // wait for Enter Key
	for i := 1; i < len(records); i++ {
		go userInput(records[i][0], done, ch)

		select {
		case <-done:
			answer := <-ch
			if records[i][1] == answer {
				correct++
			} else if answer == "Exit" {
				fmt.Println("YOU EXITED!")
				fmt.Printf("correct answers %v\nwrong answers %v\nunattempted questions %v\n", correct, incorrect, 107-(correct+incorrect))
				os.Exit(1)
			} else {
				incorrect++
			}

		case <-time.After(timeout):
			fmt.Println("Timeout!")
			fmt.Printf("correct answers %v\nwrong answers %v\nunattempted questions %v\n", correct, incorrect, 107-(correct+incorrect))
			os.Exit(1)
		}
	}

}
