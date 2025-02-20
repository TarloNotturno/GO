package mylogger

import (
	"errors"
	"fmt"
	"os"
)

type Logger interface {
	MyLog(message string)
}

func OpenFile(destinationFile string) (*os.File, error) {
	destPtr, err := os.OpenFile(destinationFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil && os.IsNotExist(err) {
		if err != nil {
			return nil, err
		}
	}
	return destPtr, nil
}

func formattingMessage(logType string, message string) string {
	return fmt.Sprintf("%s | %s\n", logType, message)
}

func CloseFile(destinationFile *os.File) error {
	return destinationFile.Close()
}

func NewLogger(input ...any) (Logger, error) {
	switch len(input) {
	case 1:
		{
			switch v := input[0].(type) {
			case string:
				return NewWindowLog(v), nil
			}
		}
	case 2:
		switch input1 := input[0].(type) {
		case string:
			{
				switch input2 := input[1].(type) {
				case *os.File:
					return NewFileLogger(input1, input2), nil
				}

			}
		}
	}
	return nil, errors.New("invalid input")
}
