package handler

import (
	"net/http"

	"github.com/raspiantoro/transaction-wrapper/internal/app/service"
)

type handlerFunc func(http.ResponseWriter, *http.Request)

type Handler struct {
	service *service.Service
	H       func(http.ResponseWriter, *http.Request)
}

func New(service *service.Service) (h *Handler) {
	h = &Handler{
		service: service,
	}
	return
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.H(w, r)
}
