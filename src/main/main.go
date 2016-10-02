package main

import (
	"fmt"
	"mouse"
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

	m.Join("#test")

	wg.Wait()
}
