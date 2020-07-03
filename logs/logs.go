package logs

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
	stdLogger   = log.New(os.Stdout, "", 0)
)

func Info(msg interface{}) {
	infoLogger.Println(msg)
}

func Error(msg string) {
	errorLogger.Println(msg)
}

func Fatal(msg string) {
	errorLogger.Println(msg)
	os.Exit(1)
}

func LogError(err error) {
	errorLogger.Println(err.Error())
}

func StdInfo(i interface{}) {
	stdLogger.Println(i)
}