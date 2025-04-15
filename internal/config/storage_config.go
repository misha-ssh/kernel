package config

import (
	"errors"
	"strings"

	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
)

const (
	FileName   = "config.txt"
	EmptyValue = ""
	Separator  = "="
)

var s *StorageConfig

var (
	ErrWriteDataToOpenFile = errors.New("write data to open file error")
	ErrCreateConfigFile    = errors.New("create config file error")
	ErrGetOpenFile         = errors.New("get open file error")
)

type StorageConfig struct {
	Storage storage.Storage
}

func (s *StorageConfig) createConfig() error {
	if !s.Storage.Exists(FileName) {
		err := s.Storage.Create(FileName)
		if err != nil {
			logger.Error(ErrCreateConfigFile)
			return ErrCreateConfigFile
		}
	}

	return nil
}

func (s *StorageConfig) validateData(key, value string) error {
	return nil
}

func Set(key, value string) error { return s.Set(key, value) }

func (s *StorageConfig) Set(key, value string) error {
	err := s.validateData(key, value)
	if err != nil {
		return err
	}

	err = s.createConfig()
	if err != nil {
		return err
	}

	param := strings.ToUpper(key) + Separator + value + "\n"

	openConfigFile, err := s.Storage.GetOpenFile(FileName)
	if err != nil {
		logger.Error(ErrGetOpenFile)
		return ErrGetOpenFile
	}

	err = s.Storage.WriteToOpenFile(openConfigFile, param)
	if err != nil {
		logger.Error(ErrWriteDataToOpenFile)
		return ErrWriteDataToOpenFile
	}

	return nil
}

func Get(key string) string { return s.Get(key) }

func (s *StorageConfig) Get(key string) string {
	got, err := s.Storage.Get(FileName)
	if err != nil {
		logger.Error(err.Error())
		return EmptyValue
	}

	return got
}

func Exists(key string) bool { return s.Exists(key) }

func (s *StorageConfig) Exists(key string) bool {
	return true
}
