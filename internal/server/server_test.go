package server

import (
	"bytes"
	"encoding/json"
	"goquizz/internal/quizz"
	"goquizz/pkg/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerQuestionsEmpty(t *testing.T) {

	q := quizz.NewQuiz([]models.Question{})
	s := NewServer(q)
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
	q := quizz.NewQuiz(questionsQuizz)
	s := NewServer(q)
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

func TestServerAnswers(t *testing.T) {

	questionsQuizz := []models.Question{
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
	q := quizz.NewQuiz(questionsQuizz)
	s := NewServer(q)

	submission := models.User{
		Username: "LeW",
		Answers: map[string]int{
			"1": 2,
			"2": 2,
		},
	}

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(submission)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/sumbit", &b)
	w := httptest.NewRecorder()
	s.HandleGetQuestions(w, req)

	res := w.Result()

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result []models.QuizResult
	if err := json.Unmarshal(body, &result); err != nil {
		return
	}

	expResult := models.QuizResult{
		TotalQuestions: 2,
		CorrectAnswers: 1,
		Percentile:     100.0,
	}

	assert.ElementsMatch(t, expResult, result)
}

func TestServerAnswers2Users(t *testing.T) {

	questionsQuizz := []models.Question{
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
	q := quizz.NewQuiz(questionsQuizz)
	s := NewServer(q)

	userW := models.User{
		Username: "LeW",
		Answers: map[string]int{
			"1": 2,
			"2": 2,
		},
	}

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(userW)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/sumbit", &b)
	w := httptest.NewRecorder()
	s.HandleGetQuestions(w, req)

	res := w.Result()

	res.Body.Close()

	expResult := models.QuizResult{
		TotalQuestions: 2,
		CorrectAnswers: 0,
		Percentile:     0.0,
	}

	userM := models.User{
		Username: "LeM",
		Answers:  map[string]int{},
	}

	err = json.NewEncoder(&b).Encode(userM)
	if err != nil {
		t.Fatal(err)
	}

	req = httptest.NewRequest(http.MethodPost, "/sumbit", &b)
	w = httptest.NewRecorder()
	s.HandleGetQuestions(w, req)

	res = w.Result()

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result []models.QuizResult
	if err := json.Unmarshal(body, &result); err != nil {
		return
	}
	assert.ElementsMatch(t, expResult, result)
}
