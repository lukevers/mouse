package javascript

import (
	"storage"
)

type Config struct {
	Name       string
	Folders    []string
	Pattern    string
	EventTypes []string
	Storage    *storage.Store
}
