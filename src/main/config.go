package main

import (
	"github.com/lukevers/viper"
	"logger/stderr"
)

type Config struct {
	Servers map[string]Server
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/mouse/")
	viper.AddConfigPath("$HOME/.mouse")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		stderr.Fatalf("Could not read config file:", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		stderr.Fatalf("Could not unmarshal config:", err)
	}
}
