package javascript

import (
	"storage"
)

// A Config represents a conriguration object for a JavaScript Plugin.
type Config struct {
	Name       string
	Folders    []string
	Pattern    string
	EventTypes []string
	Storage    *storage.Store
}
