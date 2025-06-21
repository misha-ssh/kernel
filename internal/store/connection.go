package store

import (
	"encoding/json"
	"errors"

	"github.com/ssh-connection-manager/kernel/v2/configs/envconst"
	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/crypto"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

const FileConnections = envconst.FilenameConnections

var (
	DirectionApp = storage.GetAppDir()

	ErrEncryptData   = errors.New("err encrypt data")
	ErrMarshalJson   = errors.New("failed to marshal json")
	ErrWriteJson     = errors.New("failed to write json")
	ErrUnmarshalJson = errors.New("failed to unmarshal json")
	ErrGetConnection = errors.New("failed to get connection")
	ErrDecryptData   = errors.New("failed to decrypt data")
)

func GetConnections() (*connect.Connections, error) {
	var connections connect.Connections

	cryptKey, err := GetCryptKey()
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	encryptedConnections, err := storage.Get(DirectionApp, FileConnections)
	if err != nil {
		logger.Error(ErrGetConnection.Error())
		return nil, ErrGetConnection
	}

	decryptedConnections, err := crypto.Decrypt(encryptedConnections, cryptKey)
	if err != nil {
		logger.Error(ErrDecryptData.Error())
		return nil, ErrDecryptData
	}

	err = json.Unmarshal([]byte(decryptedConnections), &connections)
	if err != nil {
		logger.Error(ErrUnmarshalJson.Error())
		return nil, ErrUnmarshalJson
	}

	return &connections, nil
}

func SetConnections(connections *connect.Connections) error {
	cryptKey, err := GetCryptKey()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

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
