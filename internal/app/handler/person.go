package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func (h *Handler) PersonHandler(w http.ResponseWriter, r *http.Request) {
	var ph handlerFunc

	switch strings.ToLower(r.Method) {
	case "post":
		ph = h.createPerson
	case "get":
		fallthrough
	default:
		ph = h.getPerson
	}

	ph(w, r)
}

func (h *Handler) getPerson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	waitCh := make(chan bool)

	go func() {
		time.Sleep(1 * time.Millisecond)
		waitCh <- true
	}()

	select {
	case <-waitCh:
		fmt.Println("Done waiting")
	case <-ctx.Done():
		fmt.Println("Request has cancelled by user")
	}

	// person, err := h.service.Person.GetPerson(ctx)
	// if err != nil {
	// 	log.Println(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }

	// data, err := json.Marshal(person)
	// if err != nil {
	// 	log.Println(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get Person Success"))
}

func (h *Handler) createPerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create Person Success"))
}
