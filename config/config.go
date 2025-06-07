package config

import (
	"encoding/json"
	"errors"
	"os/user"

	fileConfig "github.com/ssh-connection-manager/kernel/v2/internal/config"

	"github.com/ssh-connection-manager/kernel/v2/config/envconst"
	"github.com/ssh-connection-manager/kernel/v2/config/envname"
	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/crypto"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"github.com/zalando/go-keyring"
)

var (
	ErrCreateFileConnection = errors.New("err create file connection")
	ErrSetLoggerFromConfig  = errors.New("err set logger from config")
	ErrSetDefaultValue      = errors.New("err set default value")
	ErrMarshalJson          = errors.New("failed to marshal json")
	ErrWriteJson            = errors.New("failed to write json")
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
			return err
		}

		defaultConnections := &connect.Connections{
			Connects: []connect.Connect{},
		}

		jsonConnections, err := json.Marshal(defaultConnections)
		if err != nil {
			return ErrMarshalJson
		}

		currentUser, err := user.Current()
		if err != nil {
			return err
		}

		username := currentUser.Username

		cryptKey, err := keyring.Get(envconst.NameServiceCryptKey, username)
		if err != nil {
			return ErrGetCryptKey
		}

		encryptedConnections, err := crypto.Encrypt(string(jsonConnections), cryptKey)
		if err != nil {
			return ErrEncryptData
		}

		err = fileStorage.Write(filename, encryptedConnections)
		if err != nil {
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
			return ErrCreateFileConnection
		}
	}

	defaultValues := map[string]string{
		envname.Theme:  envconst.Theme,
		envname.Logger: envconst.TypeStorageLogger,
	}

	for key, value := range defaultValues {
		if !fileConfig.Exists(key) {
			err := fileConfig.Set(key, value)
			if err != nil {
				return ErrSetDefaultValue
			}
		}
	}

	return nil
}

func initCryptKey() error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	username := currentUser.Username

	service := envconst.NameServiceCryptKey

	cryptKey, _ := keyring.Get(service, username)

	if cryptKey == "" {
		cryptKey, err = crypto.GenerateKey()
		if err != nil {
			return err
		}

		err = keyring.Set(service, username, cryptKey)
		if err != nil {
			return ErrSetCryptKey
		}
	}

	return nil
}

func initLoggerFromConfig() error {
	loggerType := fileConfig.Get(envname.Logger)

	fileStorage := &storage.FileStorage{
		Direction: storage.GetHomeDir(),
	}

	switch loggerType {
	case envconst.TypeConsoleLogger:
		logger.Set(logger.NewConsoleLogger())
	case envconst.TypeStorageLogger:
		logger.Set(logger.NewStorageLogger(fileStorage))
	case envconst.TypeCombinedLogger:
		logger.Set(logger.NewCombinedLogger(
			logger.NewConsoleLogger(),
			logger.NewStorageLogger(fileStorage),
		))
	default:
		return ErrSetLoggerFromConfig
	}

	return nil
}

func Init() {
	var err error

	err = initFileConfig()
	if err != nil {
		panic(err)
	}

	err = initLoggerFromConfig()
	if err != nil {
		panic(err)
	}

	err = initCryptKey()
	if err != nil {
		panic(err)
	}

	err = initFileConnections()
	if err != nil {
		panic(err)
	}
}
