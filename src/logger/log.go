package logger

import (
	"log"
	"os"
)

var (
	Stderr = log.New(os.Stderr, "", log.LstdFlags)
	Stdout = log.New(os.Stdout, "", log.LstdFlags)
)
