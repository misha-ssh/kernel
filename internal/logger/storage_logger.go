package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
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

func (s *StorageLogger) log(value any) error {
	err := s.Storage.Create(NameLogFile)
	if err != nil {
		return ErrCreateStorage
	}

	_, calledFile, line, success := runtime.Caller(SkipUseLevel)
	if !success {
		return ErrGetInfo
	}

	logInfo := fmt.Sprintf("file: %s, line: %v, message: %#v", calledFile, line, value)

	openLogFile, err := s.Storage.GetOpenFile(NameLogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE)
	defer func(openLogFile *os.File) {
		err = openLogFile.Close()
	}(openLogFile)
	if err != nil {
		return ErrGetOpenFile
	}

	errorLog := log.New(openLogFile, "", log.LstdFlags|log.Lmicroseconds)
	errorLog.Println(logInfo)

	return nil
}

func LocStorageErr(value any) {
	localStorage := storage.LocalStorage{
		Direction: storage.HomeDir,
	}

	storageLogger := StorageLogger{
		Storage: &localStorage,
	}

	err := storageLogger.log(value)
	if err != nil {
		panic(err)
	}
}
