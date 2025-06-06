package logger

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

var ErrGetConsoleInfo = errors.New("err get info use log - console")

type ConsoleLogger struct{}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (sl *ConsoleLogger) log(value any, status StatusLog) error {
	_, calledFile, line, success := runtime.Caller(SkipUseLevel)
	if !success {
		return ErrGetConsoleInfo
	}

	logInfo := fmt.Sprintf("|%v| file: %s, line: %v, message: %#v", status, calledFile, line, value)

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println(logInfo)

	return nil
}

func (sl *ConsoleLogger) Error(value any) {
	err := sl.log(value, ErrorStatus)
	if err != nil {
		panic(err)
	}
}

func (sl *ConsoleLogger) Debug(value any) {
	err := sl.log(value, DebugStatus)
	if err != nil {
		panic(err)
	}
}

func (sl *ConsoleLogger) Info(value any) {
	err := sl.log(value, InfoStatus)
	if err != nil {
		panic(err)
	}
}

func (sl *ConsoleLogger) Warn(value any) {
	err := sl.log(value, WarnStatus)
	if err != nil {
		panic(err)
	}
}
