// Package config loads configuration for the opentimegrid server.
package config

import (
	"flag"
	"os"
)

func Server() *ServerConfiguration {
	server := new(ServerConfiguration)
	server.setFlags()
	flag.Parse()

	return server
}

func Redis() string {
	return os.Getenv("COSMOSHIPCAMP_REDIS")
}

func DB() string {
	return os.Getenv("COSMOSHIPCAMP_DB")
}
