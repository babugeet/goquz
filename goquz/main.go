package main

import (
	// "fmt"

	"github.com/babugeet/goquz/goquz/internal/quizutils"
)

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	// select {
	// case <-ctx.Done():
	// 	fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	// }

	csvfilename, timer := quizutils.GetUserArgs()
	csvContent, _ := quizutils.ReadCSVfile(csvfilename)

	quizutils.QuizQA(csvContent, timer)

}
