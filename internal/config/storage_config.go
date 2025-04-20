package config

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
)

const (
	CharNewLine = "\n"
	EmptyValue  = ""
	Separator   = "="
	FileName    = "config"
)

var (
	ErrWriteDataToOpenFile = errors.New("write data to open file error")
	ErrCreateConfigFile    = errors.New("create config file error")
	ErrGetKeyValueData     = errors.New("get value data error")
	ErrKeyOfNonLetters     = errors.New("key of non letters error")
	ErrValueIsInvalid      = errors.New("dont valid value at set data")
	ErrGetOpenFile         = errors.New("get open file error")
)

type StorageConfig struct {
	Storage storage.Storage
}

func (s *StorageConfig) create() error {
	if !s.Storage.Exists(FileName) {
		err := s.Storage.Create(FileName)
		if err != nil {
			logger.LocStorageErr(ErrCreateConfigFile)
			return ErrCreateConfigFile
		}
	}

	return nil
}

func (s *StorageConfig) rewrite(key, value string) error {
	got, err := s.Storage.Get(FileName)
	if err != nil {
		logger.LocStorageErr(err.Error())
		return err
	}

	startIndexKey := 0

	for pos, char := range got {
		if string(char) == Separator {
			if strings.ToLower(got[startIndexKey:pos]) == strings.ToLower(key) {
				neededKey := got[pos+1:]
				for i, k := range neededKey {
					if string(k) == "\n" {
						err = s.Storage.Write(FileName, got[:pos+1]+value+got[pos+i+1:])
						if err != nil {
							return err
						}

						return nil
					}
				}
			}
		}

		if string(char) == "\n" {
			startIndexKey = pos + 1
		}
	}

	return nil
}

func (s *StorageConfig) Set(key, value string) error {
	err := validateKey(key)
	if err != nil {
		return err
	}

	err = validateValue(value)
	if err != nil {
		return err
	}

	err = s.create()
	if err != nil {
		return err
	}

	if s.Exists(key) {
		err = s.rewrite(key, value)
		if err != nil {
			return err
		}

		return nil
	}

	param := strings.ToUpper(key) + Separator + value + CharNewLine

	openConfigFile, err := s.Storage.GetOpenFile(FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE)
	defer func(openConfigFile *os.File) {
		err = openConfigFile.Close()
	}(openConfigFile)
	if err != nil {
		logger.LocStorageErr(ErrGetOpenFile)
		return ErrGetOpenFile
	}

	_, err = openConfigFile.WriteString(param)
	if err != nil {
		logger.LocStorageErr(ErrWriteDataToOpenFile)
		return ErrWriteDataToOpenFile
	}

	return nil
}

func (s *StorageConfig) Get(key string) string {
	got, err := s.Storage.GetOpenFile(FileName, os.O_RDWR)
	defer func(got *os.File) {
		err = got.Close()
	}(got)
	if err != nil {
		logger.LocStorageErr(err.Error())
		return EmptyValue
	}

	sc := bufio.NewScanner(got)

	for sc.Scan() {
		data := strings.Split(sc.Text(), Separator)

		if len(data) != 2 {
			logger.LocStorageErr(ErrGetKeyValueData)
			return EmptyValue
		}

		keyConfig := data[0]
		valueConfig := data[1]

		if keyConfig == strings.ToUpper(key) {
			return valueConfig
		}
	}

	if err = sc.Err(); err != nil {
		logger.LocStorageErr(err.Error())
		return EmptyValue
	}

	return EmptyValue
}

func (s *StorageConfig) Exists(key string) bool {
	err := validateKey(key)
	if err != nil {
		logger.LocStorageErr(err.Error())
		return false
	}

	got, err := s.Storage.Get(FileName)
	if err != nil {
		logger.LocStorageErr(err.Error())
		return false
	}

	return strings.Index(got, strings.ToUpper(key)) != -1
}
