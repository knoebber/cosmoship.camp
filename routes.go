package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/knoebber/cosmoship.camp/db"
)

//go:embed client/build
var content embed.FS

func setupRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)

	staticRoutes(router)

	router.Route("/api/v1", func(api chi.Router) {
		authRoutes(api)
		memberRoutes(api)
	})
	return router
}

func staticRoutes(router *chi.Mux) {
	staticFS := fs.FS(content)

	root, err := fs.Sub(staticFS, "client/build")
	if err != nil {
		panic(fmt.Errorf("failed to get root directory for embedded content: %s", err))
	}
	fileserver := http.FileServer(http.FS(root))

	serveFile := func(w http.ResponseWriter, r *http.Request) {
		fileserver.ServeHTTP(w, r)
	}

	router.Get("/robots.txt", serveFile)
	router.Get("/favicon.ico", serveFile)
	router.Get("/asset-manifest.json", serveFile)
	router.Get("/static/*", serveFile)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if html, err := content.ReadFile("client/build/index.html"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write(html)
		}
	})
}

func authRoutes(router chi.Router) {
	router.Route("/auth", func(auth chi.Router) {
		auth.Get("/login", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("logged in"))
		})
	})
}

func memberRoutes(router chi.Router) {
	router.Route("/members", func(user chi.Router) {
		user.Post("/", func(w http.ResponseWriter, r *http.Request) {
			handleCreate(w, r, new(db.Member))
		})
		user.Get("/", func(w http.ResponseWriter, r *http.Request) {
			handleSearch(w, r, new(db.Member))
		})
		user.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			handleGet(w, r, new(db.Member))
		})
		user.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			handleDelete(w, r, new(db.Member))
		})
	})
}
