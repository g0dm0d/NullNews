package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/g0dm0d/nullnews/entity"
	"github.com/g0dm0d/nullnews/pkg/repository"
)

type NewsService struct {
	repo repository.News
}

func NewNews(repo repository.News) *NewsService {
	return &NewsService{
		repo: repo,
	}
}

func (s *NewsService) CreateNews(w http.ResponseWriter, r *http.Request) {
	var req entity.Article
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	s.repo.CreateNews(req)
}

func (s *NewsService) DeleteNews(w http.ResponseWriter, r *http.Request) {
	var req entity.Article
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	s.repo.DeleteNews(req)
}
