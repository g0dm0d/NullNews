package handler

import (
	"github.com/g0dm0d/nullnews/pkg/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Router() chi.Router {

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Get("/", h.service.CreateNews)
		r.Post("/sign-up", h.service.Register)
		r.Post("/login", h.service.Login)
	})
	return r
}
