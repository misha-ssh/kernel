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
	FileName     = "log.log"
)

var (
	ErrCreateStorage = errors.New("err at created log file")
	ErrGetOpenFile   = errors.New("err get open log file")
	ErrGetInfo       = errors.New("err get info use log")
)

type Status string

const (
	ErrorStatus Status = "ERROR"
	DebugStatus Status = "DEBUG"
	InfoStatus  Status = "INFO"
	WarnStatus  Status = "WARN"
)

var sl *StorageLogger

type StorageLogger struct {
	Storage storage.Storage
}

func init() {
	sl = New()
}

// New returns an initialized StorageLogger instance.
func New() *StorageLogger {
	localStorage := storage.LocalStorage{
		Direction: storage.GetHomeDir(),
	}

	sl := new(StorageLogger)
	sl.Storage = &localStorage

	return sl
}

func (sl *StorageLogger) createLogFile() error {
	if !sl.Storage.Exists(FileName) {
		err := sl.Storage.Create(FileName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (sl *StorageLogger) log(value any, status Status) error {
	err := sl.createLogFile()
	if err != nil {
		return ErrCreateStorage
	}

	_, calledFile, line, success := runtime.Caller(SkipUseLevel)
	if !success {
		return ErrGetInfo
	}

	logInfo := fmt.Sprintf("|%v| file: %s, line: %v, message: %#v", status, calledFile, line, value)

	openLogFile, err := sl.Storage.GetOpenFile(FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE)
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

func Error(value any) {
	err := sl.Error(value)
	if err != nil {
		panic(err)
	}
}

func (sl *StorageLogger) Error(value any) error {
	err := sl.log(value, ErrorStatus)
	if err != nil {
		return err
	}

	return nil
}

func Debug(value any) {
	err := sl.Debug(value)
	if err != nil {
		panic(err)
	}
}

func (sl *StorageLogger) Debug(value any) error {
	err := sl.log(value, DebugStatus)
	if err != nil {
		return err
	}

	return nil
}

func Info(value any) {
	err := sl.Info(value)
	if err != nil {
		panic(err)
	}
}

func (sl *StorageLogger) Info(value any) error {
	err := sl.log(value, InfoStatus)
	if err != nil {
		return err
	}

	return nil
}

func Warn(value any) {
	err := sl.Warn(value)
	if err != nil {
		panic(err)
	}
}

func (sl *StorageLogger) Warn(value any) error {
	err := sl.log(value, WarnStatus)
	if err != nil {
		return err
	}

	return nil
}
