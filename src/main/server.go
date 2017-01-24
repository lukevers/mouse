package main

import (
	"storage"
	"time"
)

type Server struct {
	Nick      string
	User      string
	Name      string
	Host      string
	Port      int
	TLS       bool
	Reconnect bool
	Ping      time.Duration
	Channels  []string
	Debug     bool
	Plugins   map[string]Plugins
	Storage   string
	Store     map[string]storage.Storage
}
