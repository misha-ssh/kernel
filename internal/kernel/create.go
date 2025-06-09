package kernel

import (
	"encoding/json"
	"errors"
	"os/user"

	"github.com/ssh-connection-manager/kernel/v2/configs/envconst"
	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/crypto"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/setup"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"github.com/zalando/go-keyring"
)

const (
	FileConnections = envconst.FilenameConnections
)

var (
	DirectionApp = storage.GetAppDir()

	ErrMarshalJson             = errors.New("failed to marshal json")
	ErrUnmarshalJson           = errors.New("failed to unmarshal json")
	ErrWriteJson               = errors.New("failed to write json")
	ErrGetCryptKey             = errors.New("err get crypt key")
	ErrGetConnection           = errors.New("failed to get connection")
	ErrDecryptData             = errors.New("failed to decrypt data")
	ErrConnectionByAliasExists = errors.New("connection by alias exists")
	ErrEncryptData             = errors.New("err encrypt data")
)

func Create(connection *connect.Connect) error {
	var connections connect.Connections

	setup.Init()

	encryptedConnections, err := storage.Get(DirectionApp, FileConnections)
	if err != nil {
		logger.Error(ErrGetConnection.Error())
		return ErrGetConnection
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

	decryptedConnections, err := crypto.Decrypt(encryptedConnections, cryptKey)
	if err != nil {
		logger.Error(ErrDecryptData.Error())
		return ErrDecryptData
	}

	err = json.Unmarshal([]byte(decryptedConnections), &connections)
	if err != nil {
		logger.Error(ErrUnmarshalJson.Error())
		return ErrUnmarshalJson
	}

	for _, savedConnection := range connections.Connects {
		if savedConnection.Alias == connection.Alias {
			logger.Error(ErrConnectionByAliasExists.Error())
			return ErrConnectionByAliasExists
		}
	}

	connections.Connects = append(connections.Connects, *connection)

	jsonConnections, err := json.Marshal(connections)
	if err != nil {
		logger.Error(ErrMarshalJson.Error())
		return ErrMarshalJson
	}

	updatedEncryptedConnections, err := crypto.Encrypt(string(jsonConnections), cryptKey)
	if err != nil {
		logger.Error(ErrEncryptData.Error())
		return ErrEncryptData
	}

	err = storage.Write(DirectionApp, FileConnections, updatedEncryptedConnections)
	if err != nil {
		logger.Error(ErrWriteJson.Error())
		return ErrWriteJson
	}

	return nil
}
