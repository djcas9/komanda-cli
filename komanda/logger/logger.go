package logger

import (
	"fmt"
	"log"
	"os"
)

var Logger *log.Logger

func Start(logPath string) {
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	Logger = log.New(f, "logs :: ", log.Lshortfile)
}
