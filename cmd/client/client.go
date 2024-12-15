package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goquizz/pkg/models"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{
		url: url,
	}
}

func (c *Client) GetQuestions() {
	resp, err := http.Get(c.url + "/questions")
	if err != nil {
		fmt.Println("Error fetching questions:", err)
		return
	}
	defer resp.Body.Close()

	var questions []models.Question
	if err := json.NewDecoder(resp.Body).Decode(&questions); err != nil {
		fmt.Println("Error decoding questions:", err)
		return
	}

	fmt.Println("Here is the quizz")
	for _, q := range questions {
		fmt.Printf("\n%s\n", q.Text)
		var sb strings.Builder
		for id, a := range q.Answers {
			fmt.Fprintf(&sb, "%s - [%s] ", id, a)
		}
		fmt.Printf("%s\n", sb.String())
	}
}

func (c *Client) TakeQuiz() {

	userAnswers := make(map[string]int)

	fmt.Print("Your username: ")
	var username string

	fmt.Scanln(&username)

	fmt.Print("Your answer (enter the question number): ")
	for i := range 2 {

		var answer int
		fmt.Printf("Question %d\n", i)
		fmt.Scanln(&answer)
		indexStr := strconv.Itoa(i)
		userAnswers[indexStr] = answer
	}

	// Submit quiz
	user := models.User{
		Answers:  userAnswers,
		Username: username,
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error preparing submission:", err)
		return
	}

	submitResp, err := http.Post(c.url+"/submit", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error submitting quiz:", err)
		return
	}
	defer submitResp.Body.Close()

	// Read and display result
	body, err := io.ReadAll(submitResp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var result models.QuizResult
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error decoding result:", err)
		return
	}

	// Display results
	fmt.Printf("\n--- Quiz Results ---\n")
	fmt.Printf("Correct Answers: %d/%d\n", result.CorrectAnswers, result.TotalQuestions)
	fmt.Printf("You performed better than %.2f%% of all quizzers\n", result.Percentile)

}
