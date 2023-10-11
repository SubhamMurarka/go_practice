package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

// func checkAnswer(answer string, correctAnswer string) (correct, incorrect int) {
// 	correct = 0
// 	incorrect = 0
// 	if answer == correctAnswer {
// 		correct++
// 	} else {
// 		incorrect++
// 	}
// 	return correct, incorrect
// }

type Pair struct {
	First  string
	Second int
}

func userInput(records [][]string, ch chan Pair) {
	for i := 1; i <= 107; i++ {
		fmt.Printf("%v\n", records[i][0])
		var answer string
		fmt.Scan(&answer)
		pair := Pair{First: answer, Second: i}
		ch <- pair
		// TODO RESET TIMER AFTER EACH INPUT
	}
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

	// fmt.Printf("%T", records)
	ch := make(chan Pair)

	var limit int = 5
	waitingTime := time.Duration(limit) * time.Second
	fmt.Printf("Time limit is %v\n", waitingTime)

	go userInput(records, ch)

	var correct int
	var incorrect int
	select {
	case input := <-ch:
		answer := input.First
		i := input.Second
		if records[i][1] == answer {
			correct++
		} else if answer == "break" {
			break
		} else {
			incorrect++
		}

	case <-time.After(waitingTime):
		fmt.Println("Timeout!")
		os.Exit(1)
	}
	fmt.Printf("correct answers %v\nwrong answers %v\nunattempted questions %v\n", correct, incorrect, 107-(correct+incorrect))
}
