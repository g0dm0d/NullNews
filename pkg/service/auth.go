package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/g0dm0d/nullnews/entity"
	"github.com/g0dm0d/nullnews/pkg/repository"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuth(repo repository.Auth) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(w http.ResponseWriter, r *http.Request) {
	var req entity.User
	json.NewDecoder(r.Body).Decode(&req)
	s.repo.Register(req)
}

func (s *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	var req entity.User
	json.NewDecoder(r.Body).Decode(&req)
	if !s.repo.Login(req) {
		log.Println("asd")
		w.WriteHeader(http.StatusUnauthorized)
	}
	// w.WriteHeader(http.StatusOK)
}
