package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from hello handler")
}
