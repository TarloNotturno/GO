package mylogger

import "fmt"

type WindowLogger struct {
	typeOfMessage string
}

func (wl WindowLogger) MyLog(message string) {
	fmt.Printf(formattingMessage(wl.typeOfMessage, message))
}

func NewWindowLog(message string) (wl WindowLogger) {
	return WindowLogger{typeOfMessage: message}
}
