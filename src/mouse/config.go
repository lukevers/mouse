package mouse

import (
	"crypto/tls"
	"storage"
)

type Config struct {
	Host          string
	Port          int
	Nick          string
	User          string
	Name          string
	Pass          string
	Reconnect     bool
	TLS           bool
	TLSConfig     *tls.Config
	Storage       storage.Storage
	StorageDriver string
}
