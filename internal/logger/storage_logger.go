package logger

import (
	"errors"
	"fmt"
	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"log"
	"runtime"
)

const (
	SkipUseLevel = 1
	NameLogFile  = "log.log"
)

var (
	ErrCreateStorage = errors.New("err at created log file")
	ErrGetOpenFile   = errors.New("err get open log file")
	ErrGetInfo       = errors.New("err get info use log")
)

type StorageLogger struct {
	Storage storage.Storage
}

func (sl *StorageLogger) log(value any) error {
	err := sl.Storage.Create(NameLogFile)
	if err != nil {
		return ErrCreateStorage
	}

	_, calledFile, line, success := runtime.Caller(SkipUseLevel)
	if !success {
		return ErrGetInfo
	}

	logInfo := fmt.Sprintf("file: %s, line: %v, message: %#v", calledFile, line, value)

	openFile, err := sl.Storage.GetOpenFile(NameLogFile)
	if err != nil {
		return ErrGetOpenFile
	}

	errorLog := log.New(openFile, "", log.LstdFlags|log.Lmicroseconds)
	errorLog.Println(logInfo)

	return nil
}

func (sl *StorageLogger) Error(value any) {
	_ = sl.log(value)
}
