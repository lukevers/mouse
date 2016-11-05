package stderr

import (
	"io"
	"logger"
)

func Fatal(v ...interface{}) {
	logger.Stderr.Fatal(v)
}

func Fatalf(format string, v ...interface{}) {
	logger.Stderr.Fatalf(format, v)
}

func Fatalln(v ...interface{}) {
	logger.Stderr.Fatalln(v)
}

func Flags() int {
	return logger.Stderr.Flags()
}

func Output(calldepth int, s string) error {
	return logger.Stderr.Output(calldepth, s)
}

func Panic(v ...interface{}) {
	logger.Stderr.Panic(v)
}

func Panicf(format string, v ...interface{}) {
	logger.Stderr.Panicf(format, v)
}

func Panicln(v ...interface{}) {
	logger.Stderr.Panicln(v)
}

func Print(v ...interface{}) {
	logger.Stderr.Print(v)
}

func Printf(format string, v ...interface{}) {
	logger.Stderr.Printf(format, v)
}

func Println(v ...interface{}) {
	logger.Stderr.Println(v)
}

func SetFlags(flag int) {
	logger.Stderr.SetFlags(flag)
}

func SetOutput(w io.Writer) {
	logger.Stderr.SetOutput(w)
}

func SetPrefix(prefix string) {
	logger.Stderr.SetPrefix(prefix)
}
