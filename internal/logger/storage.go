package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/ssh-connection-manager/kernel/v2/config/envconst"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

const FileName = envconst.FilenameLogger

var (
	DirectionApp = storage.GetAppDir()

	ErrGetStorageInfo = errors.New("err get info use log - storage")
	ErrCreateStorage  = errors.New("err at created log file")
	ErrGetOpenFile    = errors.New("err get open log file")
)

type StorageLogger struct{}

func NewStorageLogger() *StorageLogger {
	return &StorageLogger{}
}

func (sl *StorageLogger) createLogFile() error {
	if !storage.Exists(DirectionApp, FileName) {
		err := storage.Create(DirectionApp, FileName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (sl *StorageLogger) log(value any, status StatusLog) error {
	err := sl.createLogFile()
	if err != nil {
		return ErrCreateStorage
	}

	_, calledFile, line, success := runtime.Caller(SkipUseLevel)
	if !success {
		return ErrGetStorageInfo
	}

	logInfo := fmt.Sprintf("|%v| file: %s, line: %v, message: %#v", status, calledFile, line, value)

	openLogFile, err := storage.GetOpenFile(DirectionApp, FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE)
	defer func(openLogFile *os.File) {
		err = openLogFile.Close()
	}(openLogFile)
	if err != nil {
		return ErrGetOpenFile
	}

	logger := log.New(openLogFile, "", log.LstdFlags|log.Lmicroseconds)
	logger.Println(logInfo)

	return nil
}

func (sl *StorageLogger) Error(value any) {
	err := sl.log(value, ErrorStatus)
	if err != nil {
		panic(err)
	}
}

func (sl *StorageLogger) Debug(value any) {
	err := sl.log(value, DebugStatus)
	if err != nil {
		panic(err)
	}
}

func (sl *StorageLogger) Info(value any) {
	err := sl.log(value, InfoStatus)
	if err != nil {
		panic(err)
	}
}

func (sl *StorageLogger) Warn(value any) {
	err := sl.log(value, WarnStatus)
	if err != nil {
		panic(err)
	}
}
