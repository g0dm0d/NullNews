package service

import (
	"encoding/json"
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
	json.NewDecoder(r.Body).Decode(&req)
	s.repo.CreateNews(req)
}
