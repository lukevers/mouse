package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type Config struct {
	Servers map[string]Server
}

func init() {
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Println("Could not decode config.toml file")
		os.Exit(1)
	}
}
