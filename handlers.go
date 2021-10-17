package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/knoebber/cosmoship.camp/db"
	"github.com/knoebber/cosmoship.camp/models"
	"github.com/knoebber/cosmoship.camp/usererror"
)

var validate *validator.Validate

type body struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func handleGet(w http.ResponseWriter, r *http.Request, g models.Getter) {
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

func handleCreate(w http.ResponseWriter, r *http.Request, c models.Creater) {
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

func handleDelete(w http.ResponseWriter, r *http.Request, d models.Deleter) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		badRequest(w, err)
		return
	}

	if err = d.Delete(id); err != nil {
		setError(w, err)
	} else {
		setBody(w, body{Data: true})
	}
}

func handleSearch(w http.ResponseWriter, r *http.Request, s models.Searcher) {
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

func handlePasswordUpdate(w http.ResponseWriter, r *http.Request, u models.User) {
	var request struct {
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		badRequest(w, err)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		badRequest(w, err)
		return
	}

	if err := models.UpdatePassword(u, id, request.Password); err != nil {
		setError(w, err)
	} else {
		setBody(w, body{Message: "updated password"})
	}
}

func handleLoginError(w http.ResponseWriter, email string, err error) {
	log.Printf("failed login attempt for %q: %s", email, err)
	setError(w, usererror.New("email or password is incorrect"))
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

func unauthorized(w http.ResponseWriter, err error) {
	log.Printf("unauthorized: %s", err)
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

}
func permissionDenied(w http.ResponseWriter, err error) {
	log.Printf("permission denied: %s", err)
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
}

func badRequest(w http.ResponseWriter, err error) {
	log.Printf("bad request: %s", err)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func internalError(w http.ResponseWriter, err error) {
	log.Printf("internal error: %s", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
