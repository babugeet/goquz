package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/babugeet/goquz/internal/constants"
	"github.com/babugeet/goquz/internal/variables"
)

func main() {

	csvContent, _ := readCSVfile(constants.Filename)
	for _, line := range csvContent {
		fmt.Println(line[0])
		checkAnswer(getUserInput(), line[1])

	}
	fmt.Printf("Correct Answer %v , Wrong Answer %v", len(csvContent)-variables.failureCount, variables.failureCount)
}

func getUserInput() string {
	var input string
	fmt.Scan(&input)
	return input
}

func checkAnswer(userinput, actualAnswer string) {
	if userinput == actualAnswer {
		// fmt.Println("Correct Answer")
	} else {
		// fmt.Println("Wrong Answer")
		countFailure()
	}

}

func readCSVfile(filename string) ([][]string, error) {

	fmt.Println("CSV file name available is ", filename)
	fileBytes, err := os.Open(filename)
	if err != nil {
		fmt.Println("Recieved Error, ", err.Error())
		return nil, err
	}
	csvReader := csv.NewReader(fileBytes)
	csvContent, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(csvContent)
	return csvContent, nil

}

func countFailure() {
	variables.FailureCount = variables.FailureCount + 1
}
