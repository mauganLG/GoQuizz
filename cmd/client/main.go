package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	serverURL := "http://localhost:8080"

	client := NewClient(serverURL)

	quizzCmd := &cobra.Command{
		Use: "The QUIZZ",
	}
	var userNameCmd = &cobra.Command{
		Use:   "Username",
		Short: "Set up username",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Set up you username")
			var username string
			fmt.Scanln(&username)
			client.SetUpUsername(username)
		},
	}
	var questionsCmd = &cobra.Command{
		Use:   "Questions",
		Short: "Get available quiz questions",
		Run: func(cmd *cobra.Command, args []string) {
			client.GetQuestions()
		},
	}

	var takeQuizCmd = &cobra.Command{
		Use:   "Take",
		Short: "Take the quiz",
		Run: func(cmd *cobra.Command, args []string) {
			client.TakeQuiz()
		},
	}

	quizzCmd.AddCommand(userNameCmd, questionsCmd, takeQuizCmd)

	if err := quizzCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
