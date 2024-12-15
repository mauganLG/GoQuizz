package server

import (
	"encoding/json"
	"goquizz/internal/quizz"
	"goquizz/pkg/models"
	"log"
	"net/http"
)

type Server struct {
	quizz *quizz.Quiz
}

// NewServer creates a new quiz server
func NewServer(quizz *quizz.Quiz) *Server {
	return &Server{
		quizz: quizz,
	}
}

// HandleGetQuestions returns all available quiz questions
func (s *Server) HandleGetQuestions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	questions := s.quizz.GetQuestions()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

// HandleAnswers processes answers submission and returns the result
func (s *Server) HandleAnswers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Invalid submission", http.StatusBadRequest)
		return
	}

	result, err := s.quizz.SubmitAnswers(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (s *Server) HandleLenQuestions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	qLen := len(s.quizz.GetQuestions())
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(qLen)
}

// SetupRoutes link the route to their handlers
func (s *Server) SetupRoutes() {
	http.HandleFunc("/questions", s.HandleGetQuestions)
	http.HandleFunc("/submit", s.HandleAnswers)
	http.HandleFunc("/questionnumber", s.HandleLenQuestions)
}

// Start begins the HTTP server
func (s *Server) Start(port string) error {
	s.SetupRoutes()
	log.Printf("Server starting on port %s", port)
	return http.ListenAndServe(":"+port, nil)
}
