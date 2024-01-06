package core

import (
	"fmt"
	"io"
	"log"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger(f io.Writer) *Logger {
	logger := log.New(f, "", log.LstdFlags|log.Ldate|log.Llongfile)

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Print(message string) {
	l.logger.Println(message)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Printf(fmt.Sprintf(format, v...))
}

func (l *Logger) Error(err error) {
	l.logger.Printf("%+v\n", err)
}
