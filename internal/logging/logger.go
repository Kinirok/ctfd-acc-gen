package logging

import (
	"log"
)

func Init() *log.Logger {
	return log.New(log.Writer(), "[CTFd]", log.Ldate|log.Ltime|log.Lshortfile)
}
