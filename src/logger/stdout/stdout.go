package stdout

import (
	"io"
	"logger"
)

// Fatal is equivalent to the standard log package's Fatal function
func Fatal(v ...interface{}) {
	logger.Stdout.Fatal(v)
}

// Fatalf is equivalent to the standard log package's Fatalf function
func Fatalf(format string, v ...interface{}) {
	logger.Stdout.Fatalf(format, v)
}

// Fatalln is equivalent to the standard log package's Fatalln function
func Fatalln(v ...interface{}) {
	logger.Stdout.Fatalln(v)
}

// Flags is equivalent to the standard log package's Flags function
func Flags() int {
	return logger.Stdout.Flags()
}

// Output is equivalent to the standard log package's Output function
func Output(calldepth int, s string) error {
	return logger.Stdout.Output(calldepth, s)
}

// Panic is equivalent to the standard log package's Panic function
func Panic(v ...interface{}) {
	logger.Stdout.Panic(v)
}

// Panicf is equivalent to the standard log package's Panicf function
func Panicf(format string, v ...interface{}) {
	logger.Stdout.Panicf(format, v)
}

// Panicln is equivalent to the standard log package's Panicln function
func Panicln(v ...interface{}) {
	logger.Stdout.Panicln(v)
}

// Print is equivalent to the standard log package's Print function
func Print(v ...interface{}) {
	logger.Stdout.Print(v)
}

// Printf is equivalent to the standard log package's Printf function
func Printf(format string, v ...interface{}) {
	logger.Stdout.Printf(format, v)
}

// Println is equivalent to the standard log package's Println function
func Println(v ...interface{}) {
	logger.Stdout.Println(v)
}

// SetFlags is equivalent to the standard log package's SetFlags function
func SetFlags(flag int) {
	logger.Stdout.SetFlags(flag)
}

// SetOutput is equivalent to the standard log package's SetOutput function
func SetOutput(w io.Writer) {
	logger.Stdout.SetOutput(w)
}

// SetPrefix is equivalent to the standard log package's SetPrefix function
func SetPrefix(prefix string) {
	logger.Stdout.SetPrefix(prefix)
}
