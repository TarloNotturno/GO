package mylogger

import "fmt"

type WindowLogger struct {
}

func (wl WindowLogger) LogInfo(message string) {
	fmt.Printf(formattingInfo(message))
}

func (wl WindowLogger) LogError(message string) {
	fmt.Printf(formattingErr(message))
}

func NewWindowLog() (wl WindowLogger) {
	return WindowLogger{}
}
