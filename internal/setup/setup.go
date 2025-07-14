package setup

import (
	"encoding/json"
	"errors"
	"os/user"

	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/configs/envname"
	"github.com/misha-ssh/kernel/internal/config"
	"github.com/misha-ssh/kernel/internal/crypto"
	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/zalando/go-keyring"
)

var (
	ErrCreateFileConnection = errors.New("err create file connection")
	ErrSetLoggerFromConfig  = errors.New("err set logger from configs")
	ErrSetDefaultValue      = errors.New("err set default value")
	ErrMarshalJson          = errors.New("failed to marshal json")
	ErrWriteJson            = errors.New("failed to write json")
	ErrSetCryptKey          = errors.New("err set crypt key")
	ErrGetCryptKey          = errors.New("err get crypt key")
	ErrEncryptData          = errors.New("err encrypt data")
)

// initFileConnections initializes the connections file with:
// - Creates file if not exists
// - Sets default empty connections
// - Encrypts data using system keyring
// Returns error if any step fails
func initFileConnections() error {
	filename := envconst.FilenameConnections
	direction := storage.GetAppDir()

	if !storage.Exists(direction, filename) {
		err := storage.Create(direction, filename)
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

		err = storage.Write(direction, filename, encryptedConnections)
		if err != nil {
			return ErrWriteJson
		}
	}

	return nil
}

// initFileConfig initializes the config file with:
// - Creates file if not exists
// - Sets default values (theme, logger type)
// Returns error if creation or config setting fails
func initFileConfig() error {
	filename := envconst.FilenameConfig
	direction := storage.GetAppDir()

	if !storage.Exists(direction, filename) {
		err := storage.Create(direction, filename)
		if err != nil {
			return ErrCreateFileConnection
		}
	}

	defaultValues := map[string]string{
		envname.Theme:  envconst.Theme,
		envname.Logger: envconst.TypeStorageLogger,
	}

	for key, value := range defaultValues {
		if !config.Exists(key) {
			err := config.Set(key, value)
			if err != nil {
				return ErrSetDefaultValue
			}
		}
	}

	return nil
}

// initCryptKey handles encryption key setup:
// - Gets current OS user
// - Generates new key if none exists
// - Stores key in system keyring
// Returns error if key generation/storage fails
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

// initLoggerFromConfig configures logger based on:
// - Stored config value (console/storage/combined)
// Returns error if invalid logger type specified
func initLoggerFromConfig() error {
	loggerType := config.Get(envname.Logger)

	switch loggerType {
	case envconst.TypeConsoleLogger:
		logger.Set(logger.NewConsoleLogger())
	case envconst.TypeStorageLogger:
		logger.Set(logger.NewStorageLogger())
	case envconst.TypeCombinedLogger:
		logger.Set(logger.NewCombinedLogger(
			logger.NewConsoleLogger(),
			logger.NewStorageLogger(),
		))
	default:
		return ErrSetLoggerFromConfig
	}

	return nil
}

// Init performs complete application initialization:
// 1. Config file setup
// 2. Logger configuration
// 3. Crypt key initialization
// 4. Connections file setup
// Panics if any initialization fails
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
