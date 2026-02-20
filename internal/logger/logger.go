package logger

import (
	"log"
	"os"
)

func New() *log.Logger {
	return log.New(os.Stdout, "[chainproof]", log.LstdFlags|log.Lshortfile)
}
