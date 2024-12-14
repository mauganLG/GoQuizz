package server

import (
	"encoding/json"
	"fmt"
	"goquizz/internal/quizz"
	"goquizz/pkg/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerEmpty(t *testing.T) {

	q := quizz.NewQuiz([]models.Question{})
	_, err := NewServer(q)

	assert.Error(t, fmt.Errorf(""), err)
}
func TestServerQuestionsEmpty(t *testing.T) {

	q := quizz.NewQuiz([]models.Question{})
	s, _ := NewServer(q)
	req := httptest.NewRequest(http.MethodGet, "/questions", nil)

	w := httptest.NewRecorder()
	s.HandleGetQuestions(w, req)

	res := w.Result()

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var questions []models.Question
	if err := json.Unmarshal(body, &questions); err != nil {
		return
	}

	assert.ElementsMatch(t, []models.Question{}, questions)
}

func TestServerQuestions(t *testing.T) {

	questionsQuizz := []models.Question{
		{
			Id:   1,
			Text: "animal",
			Answers: []map[int]string{
				{1: "cat"},
				{2: "dog"},
			},
			CorrectAnswer: 2,
		},
		{
			Id:   2,
			Text: "letter",
			Answers: []map[int]string{
				{1: "a"},
				{2: "b"},
				{3: "c"},
			},
			CorrectAnswer: 3,
		},
	}
	q := quizz.NewQuiz(questionsQuizz)
	s, _ := NewServer(q)
	req := httptest.NewRequest(http.MethodGet, "/questions", nil)

	w := httptest.NewRecorder()
	s.HandleGetQuestions(w, req)

	res := w.Result()

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var questions []models.Question
	if err := json.Unmarshal(body, &questions); err != nil {
		return
	}

	assert.ElementsMatch(t, questionsQuizz, questions)
}
