package quizutils

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/babugeet/goquz/goquz/internal/variables"
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

func CheckAnswer(userinput, actualAnswer string) {
	if userinput == actualAnswer {
		// fmt.Println("Correct Answer")
	} else {
		// fmt.Println("Wrong Answer")
		countFailure()
	}

}

func countFailure() {
	variables.FailureCount = variables.FailureCount + 1
}

func QuizQA(csvContent [][]string) {
	for _, line := range csvContent {
		fmt.Println(line[0])
		CheckAnswer(GetUserInput(), line[1])

	}
	fmt.Printf("Correct Answer %v , Wrong Answer %v", len(csvContent)-(variables.FailureCount), (variables.FailureCount))
}

func GetUserInput() string {
	var input string
	fmt.Scan(&input)
	return input
}
