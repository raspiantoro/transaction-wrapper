package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func (h *Handler) PersonHandler(w http.ResponseWriter, r *http.Request) {
	var ph handlerFunc

	w.Header().Set("Content-Type", "application/json")

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
	var ph handlerFunc

	id := r.URL.Query().Get("id")

	switch id {
	case "":
		ph = h.getAllPerson
	default:
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", id)
		r = r.WithContext(ctx)
		ph = h.getOnePerson
	}

	ph(w, r)
}

func (h *Handler) getOnePerson(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()

	idCtx := r.Context().Value("user_id")
	id, ok := idCtx.(string)
	if !ok {
		err = errors.New("invalid userid")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	person, err := h.service.Person.GetPerson(ctx, id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	data, err := json.Marshal(person)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) getAllPerson(w http.ResponseWriter, r *http.Request) {
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

	person, err := h.service.Person.GetPersons(ctx)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	data, err := json.Marshal(person)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) createPerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create Person Success"))
}
