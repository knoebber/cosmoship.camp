package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/knoebber/cosmoship.camp/db"
	"github.com/knoebber/cosmoship.camp/usererror"
)

var validate *validator.Validate

type getter interface {
	Get(id int) error
}

type creator interface {
	Create() (interface{}, error)
}

type searcher interface {
	Search(url.Values) (interface{}, error)
}

type body struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func handleGet(w http.ResponseWriter, r *http.Request, g getter) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		badRequest(w, err)
		return
	}

	if err := g.Get(id); err != nil {
		setError(w, err)
	} else {
		setBody(w, body{Data: g})
	}
}

func handleCreate(w http.ResponseWriter, r *http.Request, c creator) {
	if err := json.NewDecoder(r.Body).Decode(c); err != nil {
		badRequest(w, err)
		return
	}

	if err := validate.Struct(c); err != nil {
		invalid(w, err)
		return
	}

	if data, err := c.Create(); err != nil {
		setError(w, err)
	} else {
		setBody(w, body{Data: data})
	}
}

func handleSearch(w http.ResponseWriter, r *http.Request, s searcher) {
	if err := r.ParseForm(); err != nil {
		badRequest(w, err)
		return
	}

	if data, err := s.Search(r.Form); err != nil {
		setError(w, err)
	} else {
		setBody(w, body{Data: data})
	}
}

func setError(w http.ResponseWriter, err error) {
	var uError *usererror.Error

	if errors.As(err, &uError) {
		customMessage(w, uError.Message)
	} else if db.NotFound(err) {
		notFound(w, err)
	} else {
		internalError(w, err)
	}
}

func invalid(w http.ResponseWriter, err error) {
	log.Printf("request invalid: %s", err)
	setBody(w, body{Message: "request invalid"})
}

func notFound(w http.ResponseWriter, err error) {
	log.Printf("resource not found: %s", err)
	setBody(w, body{Message: "resource not found"})
}

func customMessage(w http.ResponseWriter, message string) {
	log.Printf("request invalid: %s", message)
	setBody(w, body{Message: message})
}

func setBody(w http.ResponseWriter, b body) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(b); err != nil {
		internalError(w, fmt.Errorf("encoding json body: %w", err))
	}
}

func badRequest(w http.ResponseWriter, err error) {
	log.Printf("bad request: %s", err)
	http.Error(w, "bad request", http.StatusBadRequest)
}

func internalError(w http.ResponseWriter, err error) {
	log.Printf("internal error: %s", err)
	http.Error(w, "internal error", http.StatusInternalServerError)
}
