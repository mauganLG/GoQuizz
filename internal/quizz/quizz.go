package quizz

import (
	"fmt"
	"goquizz/pkg/models"
	"strconv"
)

type Quiz struct {
	questions       []models.Question
	users           []models.User
	questionsNumber models.QuestionNumber
}

// NewQuizStorage creates a new quiz storage
func NewQuiz(questions []models.Question) *Quiz {
	return &Quiz{
		questions: questions,
		users:     []models.User{},
		questionsNumber: models.QuestionNumber{
			Number: len(questions),
		},
	}
}

// GetQuestions returns all available quiz questions
func (s *Quiz) GetQuestions() []models.Question {
	return s.questions
}

// GetQuestions returns the number of question
func (s *Quiz) QuestionsNumber() models.QuestionNumber {
	return s.questionsNumber
}

// SubmitQuiz processes a quiz submission and calculates the result
func (s *Quiz) SubmitAnswers(user models.User) (models.QuizResult, error) {
	// Validate submission
	if len(user.Answers) == 0 {
		return models.QuizResult{}, fmt.Errorf("no answers submitted")
	}

	// Calculate score
	correctAnswers := 0
	for _, q := range s.questions {

		idAnswer := strconv.Itoa(q.Id)
		submittedAnswer, ok := user.Answers[idAnswer]
		if !ok {
			continue
		}
		if submittedAnswer == q.CorrectAnswer {
			correctAnswers++
		}
	}

	// Calculate percentile
	user.Score = correctAnswers
	s.users = append(s.users, user)

	// Simple percentile calculation
	percentile := s.calculatePercentile(correctAnswers)

	result := models.QuizResult{
		TotalQuestions: len(s.questions),
		CorrectAnswers: correctAnswers,
		Percentile:     percentile,
	}

	return result, nil
}

// calculatePercentile determines how the current score compares to previous submissions
func (s *Quiz) calculatePercentile(score int) float32 {
	if len(s.users) < 2 {
		return 100.0
	}

	lowerScores := 0
	for _, u := range s.users {
		if u.Score <= score {
			lowerScores++
		}
	}

	return (float32(lowerScores) / float32(len(s.users))) * 100
}
