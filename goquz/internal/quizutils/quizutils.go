package quizutils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSVfile(filename string) ([][]string, error) {

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

func CheckAnswer(userinput, actualAnswer string) bool {
	return userinput != actualAnswer
	// fmt.Println("Correct Answer")

	// fmt.Println("Wrong Answer")

}

func QuizQA(csvContent [][]string) {
	var failureCount int
	for _, line := range csvContent {
		fmt.Println(line[0])
		if !CheckAnswer(GetUserInput(), line[1]) {
			failureCount = failureCount + 1
		}

	}
	fmt.Printf("Correct Answer %v , Wrong Answer %v", (failureCount), len(csvContent)-(failureCount))
}

func GetUserInput() string {
	var input string
	fmt.Scan(&input)
	return input
}
