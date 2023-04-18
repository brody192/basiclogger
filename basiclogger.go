package basiclogger

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strings"
	"unsafe"
)

var (
	InfoBasic  *log.Logger
	Info       *log.Logger
	Warn       *log.Logger
	Error      *log.Logger
	errorBasic *log.Logger

	EchoLogger = &echoLogger{}
)

type echoLogger struct{}

func init() {
	Info = log.New(os.Stdout, "[INFO] ", log.Lshortfile)
	InfoBasic = log.New(os.Stdout, "[INFO] ", 0)
	Warn = log.New(os.Stdout, "[WARNING] ", log.Lshortfile)
	Error = log.New(os.Stdout, "[ERROR] ", log.Lshortfile)
	errorBasic = log.New(os.Stdout, "[ERROR] ", 0)
}

var lastError error

func (l *echoLogger) Error(err error) {
	err = errors.New(strings.TrimSpace(err.Error()))
	if err != lastError {
		lastError = err
		errorBasic.Println(err)
	}
}

func (l *echoLogger) Write(b []byte) (n int, err error) {
	b = bytes.TrimSpace(b)
	InfoBasic.Println(*(*string)(unsafe.Pointer(&b)))
	return len(b), nil
}
