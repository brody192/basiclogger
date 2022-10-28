package basiclogger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"unsafe"
)

var (
	InfoBasic  *log.Logger
	Info       *log.Logger
	Warn       *log.Logger
	Error      *log.Logger
	errorBasic *log.Logger

	EchoLogger = &echoLogger{}

	tz *time.Location
)

type echoLogger struct{}

type logWriter struct {
	destination io.Writer
}

func (lw *logWriter) Write(b []byte) (int, error) {
	fmt.Fprintf(lw.destination, "%s %s", time.Now().In(tz).Format("Jan-01-Mon-02-03:04.05-PM-MST"), *(*string)(unsafe.Pointer(&b)))
	return len(b), nil
}

func init() {
	Info = log.New(os.Stdout, "[INFO] ", log.Lshortfile)
	InfoBasic = log.New(os.Stdout, "[INFO] ", 0)
	Warn = log.New(os.Stdout, "[WARNING] ", log.Lshortfile)
	Error = log.New(os.Stdout, "[ERROR] ", log.Lshortfile)
	errorBasic = log.New(os.Stdout, "[ERROR] ", 0)
}

var lastError error

func (l *echoLogger) Error(err error) {
	if err != lastError {
		lastError = err
		errorBasic.Println(err)
	}
}

func (l *echoLogger) Write(b []byte) (n int, err error) {
	InfoBasic.Println(*(*string)(unsafe.Pointer(&b)))
	return len(b), nil
}

var fileOpen = false

func SaveLogFile(filename string) {
	if filename == "" {
		panic("need filename")
	}

	if !fileOpen {
		fileOpen = true

		LoadLocation("America/Toronto")

		log, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			panic(err)
		}

		var multi = io.MultiWriter(os.Stdout, &logWriter{destination: log})
		Info.SetOutput(multi)
		InfoBasic.SetOutput(multi)
		Warn.SetOutput(multi)
		Error.SetOutput(multi)
		errorBasic.SetOutput(multi)
	}
}

func LoadLocation(location string) {
	l, err := time.LoadLocation(location)
	if err != nil {
		panic(err)
	}
	tz = l
}
