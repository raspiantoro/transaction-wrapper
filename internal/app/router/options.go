package router

import (
	"github.com/raspiantoro/transaction-wrapper/internal/app/handler"
)

type Options func(r *Router)

func WithHandler(h *handler.Handler) Options {
	return func(r *Router) {
		r.handler = h
	}
}
