package main

import (
	"goquizz/internal/quizz"
	"goquizz/internal/server"
	"goquizz/pkg/models"
	"log"
)

func CreateQuestions() []models.Question {
	return []models.Question{
		{
			Id:   1,
			Text: "What animal have four legs",
			Answers: map[int]string{
				1: "cat",
				2: "bird",
				3: "snake",
				4: "fish",
			},
			CorrectAnswer: 1,
		},
		{
			Id:   2,
			Text: "What is the fourth letter of the alphabet",
			Answers: map[int]string{
				1: "a",
				2: "b",
				3: "c",
				4: "d",
			},
			CorrectAnswer: 4,
		},
		{
			Id:   3,
			Text: "What is 4x4",
			Answers: map[int]string{
				1: "3",
				2: "16",
				3: "15",
			},
			CorrectAnswer: 2,
		},
		{
			Id:   4,
			Text: "What is the largest country by population",
			Answers: map[int]string{
				1: "vatican",
				2: "Bresil",
				3: "China",
				4: "Russia",
			},
			CorrectAnswer: 3,
		},
		{
			Id:   5,
			Text: "Where is the city of Canberra",
			Answers: map[int]string{
				1: "Australia",
				2: "New Zeland",
			},
			CorrectAnswer: 2,
		},
	}

}

func main() {

	srv := server.NewServer(quizz.NewQuiz(CreateQuestions()))
	if err := srv.Start("8888"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
