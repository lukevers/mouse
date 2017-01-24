package mouse

import (
	"crypto/tls"
	"storage"
	"time"
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
	Ping          time.Duration
	TLS           bool
	TLSConfig     *tls.Config
	Storage       storage.Storage
	StorageDriver string
}
