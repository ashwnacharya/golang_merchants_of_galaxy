package logger

import (
	"os"
	"io"
	"log"
	// "io/ioutil"
)

var (
	// Trace to log just about anything
	Trace 		*log.Logger

	// Info to log important stuff
	Info  		*log.Logger 

	// Warning to log stuff to be concerned about
	Warning 	*log.Logger

	// Error to log critical stuff
	Error		*log.Logger
)

func init() {
	errorFile, err := os.OpenFile("errors.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)

	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	traceFile, err := os.OpenFile("trace.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)

	// Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Trace = log.New(io.MultiWriter(traceFile), "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(errorFile, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}