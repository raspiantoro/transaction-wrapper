package router

import (
	"net/http"

	"github.com/raspiantoro/transaction-wrapper/internal/app/handler"
)

type Router struct {
	handler *handler.Handler
}

func New(handler *handler.Handler, opts ...Options) *Router {
	r := &Router{
		handler: handler,
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r Router) Init() {
	http.HandleFunc("/person", r.handler.PersonHandler)
}
