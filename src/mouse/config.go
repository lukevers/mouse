package mouse

import (
	"crypto/tls"
	"storage"
)

// A Config represents a configuration object for a Mouse
type Config struct {
	Host          string
	Port          int
	Nick          string
	User          string
	Name          string
	Pass          string
	Channels      []string
	Reconnect     bool
	TLS           bool
	TLSConfig     *tls.Config
	Storage       storage.Storage
	StorageDriver string
}
