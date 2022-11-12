package middlewares

import (
	"net/http"

	"github.com/g0dm0d/nullnews/pkg/service"
)

type Middlewares struct {
	service *service.Service
}

func NewMid(service *service.Service) *Middlewares {
	return &Middlewares{service: service}
}

func (m *Middlewares) AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//check JWT and cookie
	})
}
