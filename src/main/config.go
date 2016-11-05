package main

import (
	"github.com/BurntSushi/toml"
	"logger/stderr"
)

type Config struct {
	Servers map[string]Server
}

func init() {
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		stderr.Fatalf("Could not decode config.toml file:", err)
	}
}
