package quizz

import (
	"goquizz/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuizz(t *testing.T) {

	q := []models.Question{
		{
			Id:   1,
			Text: "What is the capital of France?",
			Answers: map[string]string{
				"1": "London",
				"2": "Berlin",
				"3": "Paris",
				"4": "Rome",
			},
			CorrectAnswer: 3,
		},
		{
			Id:   2,
			Text: "Which programming language is this quiz built in?",
			Answers: map[string]string{
				"1": "Python",
				"2": "Go",
				"3": "JavaScript",
				"4": "Java",
			},
			CorrectAnswer: 2,
		},
	}

	quizz := NewQuiz(q)

	u := models.User{
		Username: "LeM",
		Answers: map[string]int{
			"1": 3,
			"2": 2,
		},
	}

	qr, _ := quizz.SubmitAnswers(u)

	expectedQr := models.QuizResult{
		TotalQuestions: 2,
		CorrectAnswers: 2,
		Percentile:     100,
	}

	assert.Equal(t, expectedQr, qr)
}

func TestQuizzEmpty(t *testing.T) {

	q := []models.Question{}

	quizz := NewQuiz(q)

	u := models.User{
		Username: "LeM",
		Answers: map[string]int{
			"1": 3,
			"2": 2,
		},
	}

	qr, _ := quizz.SubmitAnswers(u)

	expectedQr := models.QuizResult{
		TotalQuestions: 0,
		CorrectAnswers: 0,
		Percentile:     100,
	}

	assert.Equal(t, expectedQr, qr)
}

func TestQuizzPercentile(t *testing.T) {

	q := []models.Question{
		{
			Id:   1,
			Text: "animal",
			Answers: map[string]string{
				"1": "cat",
				"2": "dog",
			},
			CorrectAnswer: 2,
		},
		{
			Id:   2,
			Text: "letter",
			Answers: map[string]string{
				"1": "a",
				"2": "b",
				"3": "c",
			},
			CorrectAnswer: 3,
		},
	}

	quizz := NewQuiz(q)

	u := models.User{
		Username: "M",
		Answers: map[string]int{
			"1": 2,
			"2": 3,
		},
	}

	qr, _ := quizz.SubmitAnswers(u)

	u = models.User{
		Username: "W",
		Answers: map[string]int{
			"1": 1,
			"2": 3,
		},
	}

	qr, _ = quizz.SubmitAnswers(u)
	expectedQr := models.QuizResult{
		TotalQuestions: 2,
		CorrectAnswers: 1,
		Percentile:     50,
	}

	assert.Equal(t, expectedQr, qr)
}

func TestQuizzAnswerNotPresent(t *testing.T) {

	q := []models.Question{
		{
			Id:   1,
			Text: "animal",
			Answers: map[string]string{
				"1": "cat",
				"2": "dog",
			},
			CorrectAnswer: 2,
		},
		{
			Id:   2,
			Text: "letter",
			Answers: map[string]string{
				"1": "a",
				"2": "b",
				"3": "c",
			},
			CorrectAnswer: 3,
		},
	}

	quizz := NewQuiz(q)

	u := models.User{
		Username: "M",
		Answers: map[string]int{
			"1": 3,
			"2": 3,
		},
	}

	qr, _ := quizz.SubmitAnswers(u)

	expectedQr := models.QuizResult{
		TotalQuestions: 2,
		CorrectAnswers: 1,
		Percentile:     100,
	}

	assert.Equal(t, expectedQr, qr)
}
