package service

import (
	"encoding/json"
	"net/http"

	"github.com/g0dm0d/nullnews/entity"
	"github.com/g0dm0d/nullnews/pkg/repository"
)

type AuthService struct {
	repo    repository.Auth
	sectret string
}

func NewAuth(repo repository.Auth, sectret string) *AuthService {
	return &AuthService{
		repo:    repo,
		sectret: sectret,
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
	status, id := s.repo.Login(req.Password, req.Email)
	if !status {
		w.WriteHeader(http.StatusUnauthorized)
	}
	session, time := SessionGenerate()
	SetCookie(w, "token", TokenGenerate(id, s.sectret))
	SetCookie(w, "session", session)
	s.repo.SaveSession(session, id, time)
}

func SetCookie(w http.ResponseWriter, name, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: false,
	})
}

func (s *AuthService) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return
	}
	SetCookie(w, "token", "")
	SetCookie(w, "session", "")
	s.repo.DeleteSession(cookie.Value)
}
