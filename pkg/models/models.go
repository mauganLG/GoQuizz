package models

// Question represents the question structure
type Question struct {
	Id            int              `json:"id"`
	Texts         []string         `json:"texts"`
	Answers       []map[int]string `json:"answers"`
	CorrectAnswer int              `json:"correct_answer"`
}

// User represents a user's submitted quiz answers
type User struct {
	Username   string            `json:"user_id"`
	Answers    map[string]string `json:"answers"`
	Score      int               `json:"score"`
	Percentile float64           `json:"percentile"`
}

// QuizResult represents the result of a quiz submission
type QuizResult struct {
	TotalQuestions int     `json:"total_questions"`
	CorrectAnswers int     `json:"correct_answers"`
	Percentile     float64 `json:"percentile"`
}
