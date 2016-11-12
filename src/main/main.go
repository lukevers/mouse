package main

import (
	"github.com/lukevers/viper"
	"logger/stderr"
	"logger/stdout"
	"mouse"
	"mouse/plugins/scripts/javascript"
	"sync"
)

var (
	config *Config
	wg     sync.WaitGroup
	mice   []*mouse.Mouse
)

func init() {
	viper.SetConfigName("mouse")
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

func main() {
	for name, server := range config.Servers {
		m, err := mouse.New(mouse.Config{
			Host:          server.Host,
			Port:          server.Port,
			Nick:          server.Nick,
			User:          server.User,
			Name:          server.Name,
			Reconnect:     server.Reconnect,
			TLS:           server.TLS,
			Storage:       server.Store[server.Storage],
			StorageDriver: server.Storage,
		})

		if err != nil {
			stderr.Printf("Could not create mouse %s:", server.Host, err)
			break
		}

		// Display every message to STDOUT
		if server.Debug {
			m.Use(func(event *mouse.Event) {
				stdout.Println(*event)
			})
		}

		// TODO: log

		// Enable plugins if they're set to be enabled
		for language, plugin := range server.Plugins {
			switch language {
			case "javascript":
				if plugin.Enabled {
					m.Use(javascript.NewPlugin(m, &javascript.Config{
						Name:       name,
						Folders:    plugin.Folders,
						Pattern:    plugin.Pattern,
						EventTypes: plugin.Events,
						Storage:    m.Storage,
					}))
				}
			}
		}

		// Connect and join
		go func(server Server, m *mouse.Mouse) {
			// Connect
			if err := m.Connect(); err != nil {
				stderr.Printf("Could not connect to server:", err)
			}

			// Join channels
			for _, channel := range server.Channels {
				m.Join(channel)
			}
		}(server, m)

		mice = append(mice, m)
		wg.Add(1)
	}

	wg.Wait()
}
