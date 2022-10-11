package main

import (
	"log"
	"os"
)

type basicLogger struct {
	// logs [INFO] without "log.Lshortfile"
	InfoBasic *log.Logger
	Info      *log.Logger
	Warn      *log.Logger
	Error     *log.Logger
}

type echoLogger struct {
	logger basicLogger
}

var (
	Logger     = &basicLogger{}
	EchoLogger = &echoLogger{}
)

func init() {
	Logger.Info = log.New(os.Stdout, "[INFO] ", log.Lshortfile)
	Logger.InfoBasic = log.New(os.Stdout, "[INFO] ", 0)
	Logger.Warn = log.New(os.Stdout, "[WARNING] ", log.Lshortfile)
	Logger.Error = log.New(os.Stdout, "[ERROR] ", log.Lshortfile)

	EchoLogger.logger = *Logger
}

func (l *echoLogger) Error(err error) {
	l.logger.Error.Println(err)
}

func (l *echoLogger) Write(p []byte) (n int, err error) {
	l.logger.InfoBasic.Println(string(p))
	return len(p), nil
}
