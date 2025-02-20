package mylogger

import (
	"os"
)

type FileLogger struct {
	destinationFile *os.File
}

func (wl FileLogger) LogInfo(message string) {
	wl.destinationFile.WriteString(formattingInfo(message))
}

func (wl FileLogger) LogError(message string) {
	wl.destinationFile.WriteString(formattingErr(message))
}

func NewFileLogger(destinationFile *os.File) (wl FileLogger) {
	return FileLogger{destinationFile: destinationFile}
}
