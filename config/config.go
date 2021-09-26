// Package config loads configuration for the opentimegrid server.
package config

import (
	"flag"
	"os"
)

var (
	// Server stores the web server configuration.
	Server *ServerConfiguration
	DBConn string
)

// Set sets global configuration variables.
func Set() error {
	Server = new(ServerConfiguration)
	Server.setFlags()

	DBConn = os.Getenv("COSMOSHIPCAMP_DB")
	flag.Parse()

	return nil
}
