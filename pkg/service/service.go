package service

import (
	"net/http"

	"github.com/g0dm0d/nullnews/pkg/repository"
)

type News interface {
	CreateNews(http.ResponseWriter, *http.Request)
	DeleteNews(http.ResponseWriter, *http.Request)
}

type Auth interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
}

type JWT interface {
	Logout(http.ResponseWriter, *http.Request)
	RefreshJWT(http.ResponseWriter, *http.Request)
}

type Service struct {
	News
	Auth
	JWT
}

type Ctx struct {
	Secret string
}

func NewSer(rep *repository.Repository, ctx *Ctx) *Service {
	return &Service{
		News: NewNews(rep.News),
		Auth: NewAuth(rep.Auth, ctx.Secret),
		JWT:  NewJWT(rep.JWT, ctx.Secret),
	}
}
