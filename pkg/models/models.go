package models

// Question represents the question structure
type Question struct {
	Id               int            `json:"id"`
	Text             string         `json:"text"`
	AlternativeTexts []string       `json:"alternative_texts,omitempty"`
	Answers          map[int]string `json:"answers"`
	CorrectAnswer    int            `json:"correct_answer"`
}

// User represents a user's submitted quiz answers
type User struct {
	Username string      `json:"user_id"`
	Answers  map[int]int `json:"answers"`
	Score    int         `json:"score"`
}

// QuizResult represents the result of a quiz submission
type QuizResult struct {
	TotalQuestions int     `json:"total_questions"`
	CorrectAnswers int     `json:"correct_answers"`
	Percentile     float32 `json:"percentile"`
}
