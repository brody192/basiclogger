package basiclogger

import (
	"log"
	"os"
)

var (
	InfoBasic *log.Logger
	Info      *log.Logger
	Warn      *log.Logger
	Error     *log.Logger

	EchoLogger = &echoLogger{}
)

type echoLogger struct{}

func init() {
	Info = log.New(os.Stdout, "[INFO] ", log.Lshortfile)
	InfoBasic = log.New(os.Stdout, "[INFO] ", 0)
	Warn = log.New(os.Stdout, "[WARNING] ", log.Lshortfile)
	Error = log.New(os.Stdout, "[ERROR] ", log.Lshortfile)
}

func (l *echoLogger) Error(err error) {
	Error.Println(err)
}

func (l *echoLogger) Write(p []byte) (n int, err error) {
	InfoBasic.Println(string(p))
	return len(p), nil
}
