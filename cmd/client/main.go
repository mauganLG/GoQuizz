package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	serverURL := "http://localhost:8888"

	client := NewClient(serverURL)

	quizzCmd := &cobra.Command{
		Use: "The QUIZZ",
	}

	var questionsCmd = &cobra.Command{
		Use:   "questions",
		Short: "Get available quiz questions",
		Run: func(cmd *cobra.Command, args []string) {
			client.GetQuestions()
		},
	}

	var takeQuizCmd = &cobra.Command{
		Use:   "take",
		Short: "Take the quiz",
		Run: func(cmd *cobra.Command, args []string) {
			client.TakeQuiz()
		},
	}

	quizzCmd.AddCommand(questionsCmd, takeQuizCmd)

	if err := quizzCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
