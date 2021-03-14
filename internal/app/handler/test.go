package handler

import (
	"net/http"
)

func (h *Handler) ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("oops, something bad happens!"))
}
