package main

import (
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

func main() {
	for _, server := range config.Servers {
		m := mouse.New(mouse.Config{
			Host:      server.Host,
			Port:      server.Port,
			Nick:      server.Nick,
			User:      server.User,
			Name:      server.Name,
			Reconnect: server.Reconnect,
			TLS:       server.TLS,
		})

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
						Folder:     plugin.Folder,
						Pattern:    plugin.Pattern,
						EventTypes: plugin.Events,
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
				m.Join(channel, "")
			}
		}(server, m)

		mice = append(mice, m)
		wg.Add(1)
	}

	wg.Wait()
}
