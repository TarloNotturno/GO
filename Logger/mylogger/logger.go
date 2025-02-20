package mylogger

import (
	"errors"
	"fmt"
	"os"
)

type Logger interface {
	LogInfo(message string)
	LogError(message string)
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

func formattingErr(message string) string {
	return fmt.Sprintf("Error | %s\n", message)
}

func formattingInfo(message string) string {
	return fmt.Sprintf("Info | %s\n", message)
}

func CloseFile(destinationFile *os.File) error {
	return destinationFile.Close()
}

func NewLogger(input ...any) (Logger, error) {
	switch len(input) {
	case 0:
		//	{
		//		switch v := input[0].(type) {
		//		case string:
		return NewWindowLog(), nil
	//		}
	//	}
	case 1:
		//	switch input1 := input[0].(type) {
		//	case string:
		{
			switch input2 := input[0].(type) {
			case *os.File:
				return NewFileLogger(input2), nil
			}

		}
		//	}
	}
	return nil, errors.New("invalid input")
}
