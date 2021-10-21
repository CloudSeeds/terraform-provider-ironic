package logger

import (
	"fmt"
	"os"
)

func log(level, format string, v ...interface{}) {
	v = append([]interface{}{level}, v...)

	fmt.Fprintf(os.Stderr, "[%s] "+format, v...)
}

func Error(format string, v ...interface{}) {
	log("ERROR", format, v...)
}
func Warn(format string, v ...interface{}) {
	log("WARN", format, v...)
}

func Info(format string, v ...interface{}) {
	log("INFO", format, v...)
}

func Debug(format string, v ...interface{}) {
	log("DEBUG", format, v...)
}

func Trace(format string, v ...interface{}) {
	log("TRACE", format, v...)
}
