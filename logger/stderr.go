package logger

import (
	"log"
	"os"
)

func getLogger() *log.Logger {
	return log.New(os.Stderr, "", 1)
}

func writeLog(message string) {
	logger := getLogger()

	logger.Println(message)
}
