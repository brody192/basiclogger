package basiclogger

import (
	"fmt"
	"io"
	"os"
)

var Logger = &basicLogger{logger: io.Writer(os.Stdout)}

type basicLogger struct {
	logger io.Writer
}

func (l *basicLogger) Error(err error) {
	l.logger.Write([]byte(err.Error()))
}

func (l *basicLogger) Write(p []byte) (n int, err error) {
	fmt.Fprintln(l.logger, string(p))
	return len(p), nil
}
