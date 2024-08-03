package quizutils

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/babugeet/goquz/goquz/internal/constants"
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

func QuizQA(csvContent [][]string, time int) {
	var failureCount int
	// fmt.Println("changes")
	// timer1 := time.NewTimer(2 * time.Second)
	timer1 := timerStart(time)
	answerChannel := make(chan string)
	for i, line := range csvContent {
		fmt.Println(i)

		fmt.Println(line[0])

		go GetUserInput(answerChannel)
		select {
		case <-timer1.C:
			fmt.Println("Timer expired")
			fmt.Printf("Correct Answer %v , Wrong Answer %v", (failureCount), len(csvContent)-(failureCount))
			return
		case userinput := <-answerChannel:

			if !CheckAnswer(userinput, line[1]) {
				failureCount = failureCount + 1
			}
		}
	}
	fmt.Printf("Correct Answer %v , Wrong Answer %v", (failureCount), len(csvContent)-(failureCount))
	// <-channel
}

func timerStart(timeoutperiod int) *time.Timer {
	timer1 := time.NewTimer(time.Duration(timeoutperiod) * time.Second)
	return timer1
}

func GetUserInput(answerChannel chan string) {
	var input string
	fmt.Scan(&input)
	answerChannel <- input
	// return input
}

func GetUserArgs() (string, int) {
	userinput := flag.String("filepath", constants.Filename, "<filename>")
	timer := flag.Int("time", 10, "Enter time limit")
	flag.Parse()
	return *userinput, *timer
}
