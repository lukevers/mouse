package mouse

import (
	"crypto/tls"
)

type Config struct {
	Host      string
	Port      int
	Nick      string
	User      string
	Name      string
	Pass      string
	Reconnect bool
	TLS       bool
	TLSConfig *tls.Config
}
