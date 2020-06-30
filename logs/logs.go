package logs

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "info: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "error: ", log.Ldate|log.Ltime)
)

func Info(msg string) {
	infoLogger.Println(msg)
}

func Error(msg string) {
	errorLogger.Println(msg)
}

func Fatal(msg string) {
	errorLogger.Println(msg)
	os.Exit(1)
}
