package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/knoebber/cosmoship.camp/models"
)

type adminResource struct{}

func (ar adminResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(authenticate)
	r.Post("/", ar.create)
	r.Get("/", ar.search)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", ar.get)
		r.Delete("/", ar.delete)
		r.Put("/password", ar.updatePassword)
	})
	return r
}

func (ar adminResource) create(w http.ResponseWriter, r *http.Request) {
	handleCreate(w, r, new(models.Admin))
}

func (ar adminResource) get(w http.ResponseWriter, r *http.Request) {
	handleGet(w, r, new(models.Admin))
}

func (ar adminResource) delete(w http.ResponseWriter, r *http.Request) {
	handleDelete(w, r, new(models.Admin))
}

func (ar adminResource) search(w http.ResponseWriter, r *http.Request) {
	handleSearch(w, r, new(models.Admin))
}

func (ar adminResource) updatePassword(w http.ResponseWriter, r *http.Request) {
	handlePasswordUpdate(w, r, new(models.Admin))
}
