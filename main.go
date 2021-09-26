package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/knoebber/cosmoship.camp/config"
	"github.com/knoebber/cosmoship.camp/db"
)

func main() {
	if err := config.Set(); err != nil {
		panic(err)
	}

	if err := db.Start(config.DBConn); err != nil {
		panic(err)
	}
	validate = validator.New()

	timeout := time.Duration(config.Server.Timeout) * time.Second
	s := &http.Server{
		Addr:         config.Server.Addr,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	s.Handler = setupRouter()

	log.Printf("serving campers at %s", config.Server.Addr)
	log.Fatal(s.ListenAndServe())
}
