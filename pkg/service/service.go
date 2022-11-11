package service

import (
	"net/http"

	"github.com/g0dm0d/nullnews/pkg/repository"
)

type News interface {
	CreateNews(http.ResponseWriter, *http.Request)
}

type Auth interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
}

type Service struct {
	News
	Auth
}

func NewSer(rep *repository.Repository) *Service {
	return &Service{
		News: NewNews(rep.News),
		Auth: NewAuth(rep.Auth),
	}
}
