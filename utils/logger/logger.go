package logger

import (
	"log"
	"os"
	"runtime"
)

var Logger *log.Logger

func Init() {
	Logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func Info(msg string) {
	Logger.SetPrefix("INFO: ")
	logWithCallerDepth(msg, 3) // <-- adjust the depth to 3 to skip wrapper frames
}

func Error(msg string) {
	Logger.SetPrefix("ERROR: ")
	logWithCallerDepth(msg, 3)
}

func Debug(msg string) {
	Logger.SetPrefix("DEBUG: ")
	logWithCallerDepth(msg, 3)
}

func logWithCallerDepth(msg string, calldepth int) {
	// This logs msg with file:line info from calldepth stack frame
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	Logger.Printf("%s:%d: %s", file, line, msg)
}
