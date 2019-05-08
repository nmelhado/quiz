package main

// import "flag"
import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"time"
)

// Problem : the question and answer of each problem
type Problem struct {
	Question string
	Answer   string
}

func main() {
	filename := "problems.csv"

	limit := 30

	fmt.Printf("Welcome to a super easy quiz.\nPlease hit Enter to start the %v second timer and beging the quiz.\n", limit)
	fmt.Scanf("%s")
	fmt.Println("Good Luck!\n")

	timer := time.NewTicker(time.Second * time.Duration(limit))
	// Open CSV file
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read lines into a variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	var problems []Problem
	// Loop through lines and turn into object
	for _, line := range lines {
		data := Problem{
			Question: line[0],
			Answer:   line[1],
		}
		problems = append(problems, data)
	}

	correct := 0
	total := 0
	done := make(chan bool)
	go func() {
		for total < len(problems) {
			var answer string
			fmt.Println("Q: ", problems[total].Question)
			fmt.Scanf("%s", &answer)
			if answer == problems[total].Answer {
				correct++
				fmt.Println("Correct!\n")
			} else if answer == "q" {
				fmt.Println("\n\nTerminating quiz...\n")
				break
			} else {
				fmt.Println("Sorry that's wrong.")
			}
			total++
		}
		done <- true
	}()

	select {
	case <-done:
	case <-timer.C:
		fmt.Println("\n\nTime expired, terminating...\n")
	}

	var percent = math.Round(float64(correct) / float64(len(problems)) * 100)

	fmt.Printf("\nYou got %v out of %v right!  That's a %v%%!\n\n", correct, len(problems), percent)
}
