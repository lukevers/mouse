package main

import (
	"fmt"
	"mouse"
	"mouse/plugins/scripts/javascript"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	m := mouse.New(mouse.Config{
		Host:      "localhost",
		Port:      6667,
		Nick:      "mouse",
		User:      "mice",
		Name:      "mouse",
		Reconnect: true,
		TLS:       false,
	})

	wg.Add(1)

	if err := m.Connect(); err != nil {
		panic(err)
	}

	// Logger
	m.Use(func(event *mouse.Event) {
		fmt.Println(event)
		/*
			if event.Message != "" {
				fmt.Println(event.Message)
			}
		*/
	})

	m.Use(javascript.NewPlugin(m, &javascript.Config{
		Pattern:        "scripts/javascript/*.js",
		ContinuousLoad: true,
		EventTypes:     []string{"PRIVMSG"},
	}))

	m.Join("#test")

	wg.Wait()
}
