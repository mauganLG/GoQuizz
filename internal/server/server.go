package server

import (
	"encoding/json"
	"goquizz/internal/quizz"
	"net/http"
)

type Server struct {
	quizz *quizz.QuizStorage
}

// NewServer creates a new quiz server
func NewServer(quizz *quizz.QuizStorage) *Server {
	return &Server{
		quizz: quizz,
	}
}

func (s *Server) HandleGetQuestions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	questions := s.quizz.GetQuestions()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}
