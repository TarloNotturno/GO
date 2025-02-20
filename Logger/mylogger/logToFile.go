package mylogger

import (
	"os"
)

type FileLogger struct {
	typeOfMessage   string
	destinationFile *os.File
}

func (wl FileLogger) MyLog(message string) {
	wl.destinationFile.WriteString(formattingMessage(wl.typeOfMessage, message))
}

func NewFileLogger(message string, destinationFile *os.File) (wl FileLogger) {
	return FileLogger{typeOfMessage: message, destinationFile: destinationFile}
}
