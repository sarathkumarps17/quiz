package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sarathkumar17/quiz/pkg/config"

	"github.com/sarathkumar17/quiz/pkg/quiz"
	"github.com/sarathkumar17/quiz/pkg/reader"
)

type Config = config.Config

var Flags = config.Flags

func main() {

	config, err := config.GetConfig()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Please Use -help to see the help message")
		os.Exit(1)
	}
	fmt.Println("Welcome to the Quiz Game!")

	if config.Help {
		fmt.Println("Usage: quiz [flags]")
		fmt.Println("Flags:")
		for _, flag := range Flags {
			fmt.Printf("  -%s\n", flag)
			switch flag {
			case "filename":
				fmt.Println("\tThe name of the csv file containing the quiz questions and solutions defaults to 'quiz.csv'.")
			case "timeout":
				fmt.Println("\tThe maximum time allowed for the quiz in seconds defaults to 30 you can also set it to 0 to disable the timeout.")
			case "help":
				fmt.Println("\tDisplay this help message.")
			}
		}
		os.Exit(0)
	}

	if config.Filename == "" {
		config.Filename = "quiz.csv"
	}

	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	data, err := reader.ReadCSV(config.Filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(data) == 0 {
		fmt.Println("No data found in the CSV file.")
		os.Exit(1)
	}
	Quiz := quiz.MakeQuestionBank(data)
	quizChannel := make(chan string)

	// Run the quiz
	go Quiz.RunQuiz(quizChannel)

	if config.Timeout > 0 {
		// Wait for the quiz to finish on timeout
		go elapseTimeout(config.Timeout, quizChannel)
	}

	result := <-quizChannel
	if result == "Done!" {
		fmt.Println("Done!")
	}
	if result == "timeout" {
		fmt.Println("\nTime's up!")
	}
	if result == "Error!" {
		fmt.Println("\nSomething went wrong!")
	}

	Quiz.ShowResults()

}

func elapseTimeout(timeout time.Duration, quizChan chan<- string) {
	time.Sleep(timeout)
	quizChan <- "timeout"
}
