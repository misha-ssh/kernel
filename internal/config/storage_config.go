package config

import (
	"errors"
	"regexp"
	"strings"

	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
)

const (
	FileName   = "config.txt"
	EmptyValue = ""
	Separator  = "="
)

var (
	ErrWriteDataToOpenFile = errors.New("write data to open file error")
	ErrValueIsInvalid      = errors.New("dont valid value at set data")
	ErrCreateConfigFile    = errors.New("create config file error")
	ErrKeyOfNonLetters     = errors.New("key of non letters error")
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
	matchedKey, err := regexp.MatchString("^[a-zA-Z]", key)
	if err != nil {
		return err
	}

	if !matchedKey {
		return ErrKeyOfNonLetters
	}

	if strings.TrimSpace(value) == "" {
		return ErrValueIsInvalid
	}

	return nil
}

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

func (s *StorageConfig) Get(key string) string {
	got, err := s.Storage.Get(FileName)
	if err != nil {
		logger.Error(err.Error())
		return EmptyValue
	}

	startIndexKey := 0

	for pos, char := range got {
		if string(char) == Separator {
			if strings.ToLower(got[startIndexKey:pos]) == strings.ToLower(key) {
				neededKey := got[pos+1:]
				for i, k := range neededKey {
					if string(k) == "\n" {
						return neededKey[:i]
					}
				}
			}
		}

		if string(char) == "\n" {
			startIndexKey = pos + 1
		}
	}

	return EmptyValue
}

func (s *StorageConfig) Exists(key string) bool {
	return true
}
