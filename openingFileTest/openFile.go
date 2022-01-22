package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// pass filepath
	filepath := "../problems1.csv"

	//open file
	openfile, err := os.Open(filepath)
	checkError("Error in opening file\n", err)
	
	// read file
	filedata, err := csv.NewReader(openfile).ReadAll()
	checkError("Error in reading file\n", err)
	for _, val := range filedata {
		fmt.Println(val)
	}
}

func checkError(msg string, err error) {
	if err != nil {
		log.Fatalf(msg,  err)
	}
}