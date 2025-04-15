package config

import (
	"os"
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

type StorageConfig struct {
	Storage storage.Storage
}

func (s *StorageConfig) createConfig() {
	if !s.Storage.Exists(FileName) {
		err := s.Storage.Create(FileName)
		if err != nil {
			logger.Error(err.Error())
		}
	}
}

func (s *StorageConfig) validateData()

func Set(key, value string) { s.Set(key, value) }

func (s *StorageConfig) Set(key, value string) {
	s.createConfig()

	param := strings.ToUpper(key) + Separator + value + "\n"

	openConfigFile, err := s.Storage.GetOpenFile(FileName)
	defer func(openConfigFile *os.File) {
		err = openConfigFile.Close()
	}(openConfigFile)
	if err != nil {
		logger.Error(err.Error())
	}

	_, err = openConfigFile.WriteString(param)
	if err != nil {
		logger.Error(err.Error())
	}
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
