package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/knoebber/cosmoship.camp/config"
	"github.com/knoebber/cosmoship.camp/db"
	"github.com/knoebber/cosmoship.camp/models"
	"github.com/knoebber/cosmoship.camp/redispool"
)

const timeout = 20 * time.Second

func main() {
	serverConfig := config.Server()
	validate = validator.New()

	if err := setup(); err != nil {
		log.Fatalf("setting up server: %s", err)
	}

	s := &http.Server{
		Addr:         serverConfig.Addr,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	s.Handler = setupRouter(serverConfig)

	log.Printf("serving campers at %s", serverConfig.Addr)
	log.Fatal(s.ListenAndServe())
}

func setup() error {
	if err := redispool.Start(config.Redis()); err != nil {
		return err
	}
	if err := db.Start(config.DB()); err != nil {
		return err
	}
	if err := models.Migrate(); err != nil {
		return err
	}

	return nil
}
