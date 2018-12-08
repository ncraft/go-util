// Very simple logging implementation with log levels DEBUG, INFO and ERROR.
package log

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var (
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	debug bool
	once  sync.Once
)

func SetDebug(value bool) {
	once.Do(func() {
		debug = value
	})
}

func Debug(format string, args ...interface{}) {
	if debug {
		debugLogger.Output(2, fmt.Sprintf(format, args...))
	}
}

func Info(format string, args ...interface{}) {
	infoLogger.Output(2, fmt.Sprintf(format, args...))
}

func InfoFunc(funcName string, format string, args ...interface{}) {
	// function name could as well be retrieved in runtime: https://wycd.net/posts/2014-07-02-logging-function-names-in-go.html
	infoLogger.Output(2, fmt.Sprintf(funcName+": "+format, args...))
}

func Error(format string, args ...interface{}) {
	errorLogger.Output(2, fmt.Sprintf(format, args...))
}

func ErrorDetail(message string, err error) {
	errorLogger.Output(2, fmt.Sprintf("%s: %+v", message, err))
}

func ErrorJustDetail(err error) {
	errorLogger.Output(2, fmt.Sprintf("%+v", err))
}
