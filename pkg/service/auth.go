package service

import (
	"encoding/json"
	"log"
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
	status, user := s.repo.Login(req.Password, req.Email)
	if !status {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	session, time := SessionGenerate()
	sessionID, err := s.repo.SaveSession(session, user.ID, time)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	b, err := json.Marshal(&Session{Session: session})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token := TokenGenerate(user, sessionID, s.sectret)
	SetCookie(w, "token", token)
	w.Write([]byte(b))
}
