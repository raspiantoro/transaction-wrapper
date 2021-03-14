package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/raspiantoro/transaction-wrapper/internal/app/payload"
	"gorm.io/gorm"
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
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	person, err := h.service.Person.GetPerson(ctx, id)
	if err == gorm.ErrRecordNotFound {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("oops...something bad happen"))
		return
	}

	data, err := json.Marshal(person)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
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
	if err == gorm.ErrRecordNotFound {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("oops...something bad happen"))
		return
	}

	data, err := json.Marshal(person)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h *Handler) createPerson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var req payload.CreatePersonRequests
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// change to CreatePerson, to test without transaction
	err = h.service.Person.CreatePersonTx(ctx, req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("oops...something bad happen"))
		return
	}

	data := payload.CreatePersonResponses{
		Message: "Create Person Success",
	}

	resp, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("oops...something bad happen"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
