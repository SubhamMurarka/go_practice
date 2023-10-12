package main

import (
	"encoding/csv"
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

func main() {

	fd, error := os.Open("/home/murarka/go_project/problems.csv")

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Successfully opened the csv file")
	defer fd.Close()

	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()

	if error != nil {
		log.Fatal(error)
	}

	timeout := 5 * time.Second
	ch := make(chan string)
	var correct int
	var incorrect int

	for i := 1; i <= 107; i++ {
		done := make(chan bool)
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
