package config

import "flag"

// ServerConfiguration configures the web server.
type ServerConfiguration struct {
	Addr    string `json:"addr"`
	Timeout int    `json:"timeout"`
}

func (s *ServerConfiguration) setFlags() {
	flag.StringVar(&s.Addr, "addr", ":3001", "http address to listen on")
	flag.IntVar(&s.Timeout, "timeout", 10, "http server timeout in seconds")
}
