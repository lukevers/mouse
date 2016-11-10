package main

type Server struct {
	Nick      string
	User      string
	Name      string
	Host      string
	Port      int
	TLS       bool
	Reconnect bool
	Channels  []string
	Debug     bool
	Plugins   map[string]struct {
		Enabled bool
		Folders []string
		Pattern string
		Events  []string
	}
}
