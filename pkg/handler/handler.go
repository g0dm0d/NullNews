package handler

import (
	"net/http"

	"github.com/g0dm0d/nullnews/pkg/middlewares"
	"github.com/g0dm0d/nullnews/pkg/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service     *service.Service
	middlewares *middlewares.Middlewares
}

func New(service *service.Service, middlewares *middlewares.Middlewares) *Handler {
	return &Handler{
		service:     service,
		middlewares: middlewares,
	}
}

func (h *Handler) Router() chi.Router {

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Post("/sign-up", h.service.Register)
		r.Post("/login", h.service.Login)
		r.Post("/delete-article", h.service.DeleteNews)
		r.Post("/refresh-jwt", h.service.RefreshJWT)
	})

	r.Group(func(r chi.Router) {
		r.Use(h.middlewares.User)
		r.Post("/logout", h.service.Logout)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!"))
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(h.middlewares.Writer)
		r.Post("/new-article", h.service.CreateNews)
	})
	return r
}
