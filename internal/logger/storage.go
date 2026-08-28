package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/internal/storage"
)

const Filename = envconst.FilenameLogger

var (
	ErrGetStorageInfo = errors.New("err get info use log - storage")
	ErrCreateStorage  = errors.New("err at created log file")
	ErrGetOpenFile    = errors.New("err get open log file")
)

type StorageLogger struct {
	storage *storage.Local
}

func NewStorageLogger() *StorageLogger {
	return &StorageLogger{
		storage: storage.NewLocal(),
	}
}

func (s *StorageLogger) createLogFile() error {
	if !s.storage.Exists(Filename) {
		err := s.storage.Create(Filename)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StorageLogger) log(value any, status StatusLog) error {
	err := s.createLogFile()
	if err != nil {
		return ErrCreateStorage
	}

	_, calledFile, line, success := runtime.Caller(SkipUseLevel)
	if !success {
		return ErrGetStorageInfo
	}

	logInfo := fmt.Sprintf("|%v| file: %s, line: %v, message: %#v", status, calledFile, line, value)

	openLogFile, err := s.storage.GetOpenFile(Filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE)
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

func (s *StorageLogger) Error(value any) {
	err := s.log(value, ErrorStatus)
	if err != nil {
		panic(err)
	}
}

func (s *StorageLogger) Debug(value any) {
	err := s.log(value, DebugStatus)
	if err != nil {
		panic(err)
	}
}

func (s *StorageLogger) Info(value any) {
	err := s.log(value, InfoStatus)
	if err != nil {
		panic(err)
	}
}

func (s *StorageLogger) Warn(value any) {
	err := s.log(value, WarnStatus)
	if err != nil {
		panic(err)
	}
}
