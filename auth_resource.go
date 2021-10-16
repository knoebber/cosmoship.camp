package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/knoebber/cosmoship.camp/models"
)

const sessionCookie = "cosmoshipcamp-session"

func init() {
	rand.Seed(time.Now().UnixNano())
}

type authResource struct{ secure bool }

func (ar authResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/login", ar.login())
	r.Get("/session", ar.getSession)
	return r
}

func (ar authResource) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		// Randomization the response time.
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(3000)))

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			badRequest(w, err)
			return
		}

		session, err := models.Login(request.Email, request.Password, r.RemoteAddr)
		if err != nil {
			handleLoginError(w, request.Email, err)
			return
		}

		http.SetCookie(w, &http.Cookie{
			HttpOnly: true,
			MaxAge:   models.SessionDurationSeconds,
			Name:     sessionCookie,
			SameSite: http.SameSiteStrictMode,
			Secure:   ar.secure,
			Value:    session,
		})
		setBody(w, body{Message: "logged in"})
	}
}

func (ar authResource) getSession(w http.ResponseWriter, r *http.Request) {
	var g getter
	id, isAdmin := checkSession(w, r)
	if id < 1 {
		return
	}
	if isAdmin {
		g = new(models.Admin)
	} else {
		g = new(models.Member)
	}
	if err := g.Get(id); err != nil {
		setError(w, err)
	} else {
		setBody(w, body{Data: g})
	}
}

func checkSession(w http.ResponseWriter, r *http.Request) (id int, isAdmin bool) {
	sessionKey, err := r.Cookie(sessionCookie)
	if err != nil {
		unauthorized(w, err)
		return
	}

	sessionValue, err := models.CheckSession(sessionKey.String())
	if err != nil {
		// Unset the session cookie.
		http.SetCookie(w, &http.Cookie{Name: sessionCookie, MaxAge: -1})
		unauthorized(w, err)
		return
	}

	return models.ParseSessionValue(sessionValue)
}
