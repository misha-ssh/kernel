package config

import (
	"errors"
	"os/user"

	"github.com/ssh-connection-manager/kernel/v2/config/envconst"
	"github.com/ssh-connection-manager/kernel/v2/config/envname"
	"github.com/ssh-connection-manager/kernel/v2/internal/config"
	"github.com/ssh-connection-manager/kernel/v2/internal/crypto"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"github.com/zalando/go-keyring"
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

// todo добавить логику при которой будет хешироваться ключ
// чтобы можно было его потом расшироваться (более безопаснее будет данный метод)
func initCryptKey() error {
	currentUser, err := user.Current()
	if err != nil {
		logger.Error(err)
		return err
	}

	username := currentUser.Username
	service := envconst.NameServiceCryptKey

	cryptKey, _ := keyring.Get(service, username)

	if len(cryptKey) == 0 {
		generatedKey, err := crypto.GenerateKey()
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		err = keyring.Set(service, username, generatedKey)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
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

	err = initCryptKey()
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
