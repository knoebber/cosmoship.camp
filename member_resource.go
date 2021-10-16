package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/knoebber/cosmoship.camp/models"
)

type memberResource struct{}

func (mr memberResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(authenticate)
	r.Post("/", mr.create)
	r.Get("/", mr.search)
	r.Get("/sources", mr.sources)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", mr.get)
		r.Delete("/", mr.delete)
		r.Put("/password", mr.updatePassword)
	})
	return r
}

func (mr memberResource) create(w http.ResponseWriter, r *http.Request) {
	handleCreate(w, r, new(models.Member))
}

func (mr memberResource) get(w http.ResponseWriter, r *http.Request) {
	handleGet(w, r, new(models.Member))
}

func (mr memberResource) delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, new(models.Member))
}

func (mr memberResource) search(w http.ResponseWriter, r *http.Request) {
	handleSearch(w, r, new(models.Member))
}

func (mr memberResource) sources(w http.ResponseWriter, r *http.Request) {
	setBody(w, body{
		Data: []models.MemberSource{
			models.MemberSourcePastBooking,
			models.MemberSourceGuestBook,
			models.MemberSourceInPerson,
		},
	})
}

func (mr memberResource) updatePassword(w http.ResponseWriter, r *http.Request) {
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

	if err := models.UpdateMemberPassword(id, request.Password); err != nil {
		setError(w, err)
	} else {
		setBody(w, body{Message: "updated password"})
	}
}
