package middlewares

import (
	"net/http"

	"github.com/g0dm0d/nullnews/pkg/service"
)

type Middlewares struct {
	service *service.Service
	ctx     *Ctx
}

type Ctx struct {
	Secret string
}

func NewMid(service *service.Service, ctx *Ctx) *Middlewares {
	return &Middlewares{
		service: service,
		ctx:     ctx,
	}
}

func (m *Middlewares) AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			return
		}
		_, err = service.TokenParse(token.Value, m.ctx.Secret)
		if err != nil {
			return
		}
		next.ServeHTTP(w, r)
	})
}
