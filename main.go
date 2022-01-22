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

	//naming the file and adding a help tag
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer' ") 
	timeLimit := flag.Int("limit", 30, "time limit for the quiz in seconds")
	flag.Parse()
	// _ = csvFilename // Only for check

	// Opening the file
	file, err := os.Open(*csvFilename)
		checkError("Error in opening file: %s\n", *csvFilename, err)
	
	// Read file
// 	filedata, err := csv.NewReader(openfile).ReadAll()
// 	checkError("Error in reading file: \n", err)
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	checkError("Error reading file: %s\n", *csvFilename, err)
	
	problems := parseLine(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	
	// Print the problems to the user
	correct := 0
	for i, p := range problems {
					fmt.Printf("Problem #%d: %s = ", i+1, p.q)
					answerCh := make(chan string)
					go func () {
						var answer string
						fmt.Scanf("%s\n", &answer)
						answerCh <- answer
					} ()
		select {
		case <- timer.C:
				fmt.Print("\nYou are out of time. See result below:")
				fmt.Printf("\n%d correct answers out of %d questions\n", correct, len(problems) )
				return
		case answer := <- answerCh:
		if answer == p.a {
			correct++
			fmt.Println("You are correct!")
		} else {
			fmt.Println("Wrong answer!")}		
		}
		
	}
	fmt.Printf("%d correct answers out of %d questions\n", correct, len(problems) )
}

func parseLine(lines [] [] string) [] problem {
	ret := make([] problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			// Since we know the exact size of our slice
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
		return ret
}

type problem struct {
	q string
	a string
}

func checkError(msg string, fileName string, err error) {
	if err != nil {
		log.Fatal(msg, fileName, err)
		os.Exit(1)
	} 	
}