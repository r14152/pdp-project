package logger

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota + 1
)

var (
	l        *log.Logger // PRIVATE LOGGER INSTANCE
	logLevel LogLevel    // PRIVATE LOGGING LEVEL
)

type LoggerConfig struct {
	logFile  string
	logLevel int
}

func Init(config *LoggerConfig) bool {
	fileHandle := os.Stderr
	l = log.New(fileHandle, "", log.Ldate|log.Ltime|log.Lshortfile)
	l.Println("Logger initialised successfully.")
	//TODO
	return true
}

func llog(format string, v ...interface{}) {
	l.Output(3, fmt.Sprintf(format, v...))
}

func Debug(format string, v ...interface{}) {
	llog(format, v)
	//TODO
}

func SetLogLevel(level LogLevel) {
	//TODO
}
