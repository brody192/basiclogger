package basiclogger

import (
	"log"
	"os"
)

var (
	Info       = log.New(os.Stdout, "[INFO] ", log.Lshortfile)
	InfoBasic  = log.New(os.Stdout, "[INFO] ", 0)
	Warn       = log.New(os.Stdout, "[WARNING] ", log.Lshortfile)
	Error      = log.New(os.Stderr, "[ERROR] ", log.Lshortfile)
	ErrorBasic = log.New(os.Stderr, "[ERROR] ", 0)
)
