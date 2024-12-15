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
		t.Fatal(err)
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
		t.Fatal(err)
	}

	assert.ElementsMatch(t, questionsQuizz, questions)
}

func TestServerLenQuestions(t *testing.T) {

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
		{
			Id:   3,
			Text: "malcom",
			Answers: map[string]string{
				"1": "yes",
				"2": "no",
				"3": "maybe",
			},
			CorrectAnswer: 3,
		},
	}
	q := quizz.NewQuiz(questionsQuizz)
	s := NewServer(q)
	req := httptest.NewRequest(http.MethodGet, "/questionnumber", nil)

	w := httptest.NewRecorder()
	s.HandleLenQuestions(w, req)

	res := w.Result()

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var questionsNumber models.QuestionNumber
	if err := json.Unmarshal(body, &questionsNumber); err != nil {
		t.Fatal(err)
	}

	qn := models.QuestionNumber{
		Number: 3,
	}

	assert.Equal(t, qn, questionsNumber)
}

func TestServerLenQuestionsEmpty(t *testing.T) {

	questionsQuizz := []models.Question{}
	q := quizz.NewQuiz(questionsQuizz)
	s := NewServer(q)
	req := httptest.NewRequest(http.MethodGet, "/questionnumber", nil)

	w := httptest.NewRecorder()
	s.HandleLenQuestions(w, req)

	res := w.Result()

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var questionsNumber models.QuestionNumber
	if err := json.Unmarshal(body, &questionsNumber); err != nil {
		t.Fatal(err)
	}

	qn := models.QuestionNumber{
		Number: 0,
	}
	assert.Equal(t, qn, questionsNumber)
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

	req := httptest.NewRequest(http.MethodPost, "/submit", &b)
	w := httptest.NewRecorder()
	s.HandleGetQuestions(w, req)

	res := w.Result()

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result models.QuizResult
	if err := json.Unmarshal(body, &result); err != nil {
		return
	}

	expResult := models.QuizResult{
		TotalQuestions: 2,
		CorrectAnswers: 1,
		Percentile:     100.0,
	}
	assert.Equal(t, expResult, result)
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

	req := httptest.NewRequest(http.MethodPost, "/submit", &b)
	w := httptest.NewRecorder()
	s.HandleGetQuestions(w, req)

	res := w.Result()

	res.Body.Close()

	userM := models.User{
		Username: "MLG",
		Answers:  map[string]int{},
	}

	err = json.NewEncoder(&b).Encode(userM)
	if err != nil {
		t.Fatal(err)
	}

	req = httptest.NewRequest(http.MethodPost, "/submit", &b)
	w = httptest.NewRecorder()
	s.HandleGetQuestions(w, req)

	res = w.Result()

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var result models.QuizResult
	if err := json.Unmarshal(body, &result); err != nil {
		return
	}

	expResult := models.QuizResult{
		TotalQuestions: 2,
		CorrectAnswers: 0,
		Percentile:     0.0,
	}

	assert.Equal(t, expResult, result)
}
