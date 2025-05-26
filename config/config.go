package config

import (
	"errors"

	"github.com/ssh-connection-manager/kernel/v2/config/envconst"
	"github.com/ssh-connection-manager/kernel/v2/config/envname"
	"github.com/ssh-connection-manager/kernel/v2/internal/config"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

var ErrGetConsoleInfo = errors.New("err set default value")

func setDefaultValues(config config.Config) error {
	if !config.Exists(envname.Theme) {
		err := config.Set(envname.Theme, envconst.Theme)
		if err != nil {
			logger.Error(ErrGetConsoleInfo)
			return ErrGetConsoleInfo
		}
	}

	return nil
}

// todo добавить дефолтные значения при пустом файле и так же их шифрануть
func initFileConnections() error {
	fileName := envconst.FilenameConnection

	fileStorage := &storage.FileStorage{
		Direction: storage.GetHomeDir(),
	}

	if !fileStorage.Exists(fileName) {
		err := fileStorage.Create(envconst.FilenameConnection)
		if err != nil {
			return err
		}
	}

	return nil
}

func initFileConfig() error {
	fileStorage := &storage.FileStorage{
		Direction: storage.GetHomeDir(),
	}

	fileConfig := &config.StorageConfig{
		Storage: fileStorage,
	}

	err := setDefaultValues(fileConfig)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func Init() error {
	err := initFileConfig()
	if err != nil {
		logger.Error(err)
		return err
	}

	err = initFileConnections()
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
