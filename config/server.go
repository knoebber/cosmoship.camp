package config

import "flag"

// ServerConfiguration configures the web server.
type ServerConfiguration struct {
	Addr   string
	Secure bool
}

func (s *ServerConfiguration) setFlags() {
	flag.StringVar(&s.Addr, "addr", ":3001", "http address to listen on")
	flag.BoolVar(&s.Secure, "secure", false, "set true when using https")
}
