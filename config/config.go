package config

import (
	"encoding/json"
	"errors"
	"os/user"

	"github.com/ssh-connection-manager/kernel/v2/config/envconst"
	"github.com/ssh-connection-manager/kernel/v2/config/envname"
	"github.com/ssh-connection-manager/kernel/v2/internal/config"
	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/crypto"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"github.com/zalando/go-keyring"
)

var (
	ErrMarshalJson          = errors.New("failed to marshal json")
	ErrGetConsoleInfo       = errors.New("err set default value")
	ErrWriteJson            = errors.New("failed to write json")
	ErrCreateFileConnection = errors.New("err create file connection")
	ErrSetCryptKey          = errors.New("err set crypt key")
	ErrGetCryptKey          = errors.New("err get crypt key")
	ErrEncryptData          = errors.New("err encrypt data")
)

func initFileConnections() error {
	filename := envconst.FilenameConnection

	fileStorage := &storage.FileStorage{
		Direction: storage.GetHomeDir(),
	}

	if !fileStorage.Exists(filename) {
		err := fileStorage.Create(filename)
		if err != nil {
			logger.Error(ErrCreateFileConnection.Error())
			return err
		}

		defaultConnections := &connect.Connections{
			Connects: []connect.Connect{},
		}

		jsonConnections, err := json.Marshal(defaultConnections)
		if err != nil {
			logger.Error(ErrMarshalJson.Error())
			return ErrMarshalJson
		}

		currentUser, err := user.Current()
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		username := currentUser.Username

		cryptKey, err := keyring.Get(envconst.NameServiceCryptKey, username)
		if err != nil {
			logger.Error(ErrGetCryptKey.Error())
			return err
		}

		encryptedConnections, err := crypto.Encrypt(string(jsonConnections), cryptKey)
		if err != nil {
			logger.Error(ErrEncryptData.Error())
			return ErrEncryptData
		}

		err = fileStorage.Write(filename, encryptedConnections)
		if err != nil {
			logger.Error(ErrWriteJson.Error())
			return ErrWriteJson
		}
	}

	return nil
}

func initFileConfig() error {
	filename := envconst.FilenameConfig

	fileStorage := &storage.FileStorage{
		Direction: storage.GetHomeDir(),
	}

	if !fileStorage.Exists(filename) {
		err := fileStorage.Create(filename)
		if err != nil {
			logger.Error(ErrCreateFileConnection.Error())
			return err
		}
	}

	fileConfig := &config.StorageConfig{
		Storage: fileStorage,
	}

	if !fileConfig.Exists(envname.Theme) {
		err := fileConfig.Set(envname.Theme, envconst.Theme)
		if err != nil {
			logger.Error(ErrGetConsoleInfo.Error())
			return ErrGetConsoleInfo
		}
	}

	return nil
}

func initCryptKey() error {
	currentUser, err := user.Current()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	username := currentUser.Username

	service := envconst.NameServiceCryptKey

	cryptKey, _ := keyring.Get(service, username)

	if cryptKey == "" {
		cryptKey, err = crypto.GenerateKey()
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		err = keyring.Set(service, username, cryptKey)
		if err != nil {
			logger.Error(ErrSetCryptKey.Error())
			return ErrSetCryptKey
		}
	}

	return nil
}

func Init() error {
	var err error

	err = initCryptKey()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = initFileConfig()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = initFileConnections()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
