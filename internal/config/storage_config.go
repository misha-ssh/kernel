package config

import (
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"strings"
)

const (
	FileName = "config.txt"
)

var l *StorageConfig

type StorageConfig struct {
	Storage storage.Storage
}

func Get(key string) string { return l.Get(key) }

func (l *StorageConfig) Get(key string) string {
	return ""
}

func Set(key, value string) { l.Set(key, value) }

func (l *StorageConfig) Set(key, value string) {
	if !l.Storage.Exists(FileName) {
		err := l.Storage.Create(FileName)
		if err != nil {
			logger.Error(err.Error())
		}
	}

	param := strings.ToUpper(key) + "=" + value

	err := l.Storage.Write(FileName, param)
	if err != nil {
		logger.Error(err.Error())
	}
}

func Exists(key string) bool { return l.Exists(key) }

func (l *StorageConfig) Exists(key string) bool {
	return true
}
