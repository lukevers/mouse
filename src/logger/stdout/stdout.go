package stdout

import (
	"io"
	"logger"
)

func Fatal(v ...interface{}) {
	logger.Stdout.Fatal(v)
}

func Fatalf(format string, v ...interface{}) {
	logger.Stdout.Fatalf(format, v)
}

func Fatalln(v ...interface{}) {
	logger.Stdout.Fatalln(v)
}

func Flags() int {
	return logger.Stdout.Flags()
}

func Output(calldepth int, s string) error {
	return logger.Stdout.Output(calldepth, s)
}

func Panic(v ...interface{}) {
	logger.Stdout.Panic(v)
}

func Panicf(format string, v ...interface{}) {
	logger.Stdout.Panicf(format, v)
}

func Panicln(v ...interface{}) {
	logger.Stdout.Panicln(v)
}

func Print(v ...interface{}) {
	logger.Stdout.Print(v)
}

func Printf(format string, v ...interface{}) {
	logger.Stdout.Printf(format, v)
}

func Println(v ...interface{}) {
	logger.Stdout.Println(v)
}

func SetFlags(flag int) {
	logger.Stdout.SetFlags(flag)
}

func SetOutput(w io.Writer) {
	logger.Stdout.SetOutput(w)
}

func SetPrefix(prefix string) {
	logger.Stdout.SetPrefix(prefix)
}
