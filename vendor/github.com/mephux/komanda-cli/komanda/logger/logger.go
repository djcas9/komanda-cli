package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger global
// TODO: fix later - bad
var Logger *log.Logger

// Start logging komanda-cli information to the default log location
func Start(logPath string) {
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	Logger = log.New(f, "logs :: ", log.Lshortfile)
}
