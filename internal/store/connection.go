package store

import (
	"encoding/json"
	"errors"

	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/internal/crypto"
	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
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

// GetConnections get connection from file
func GetConnections() (*connect.Connections, error) {
	var connections connect.Connections

	cryptKey, err := GetCryptKey()
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	encryptedConnections, err := storage.Get(DirectionApp, FileConnections)
	if err != nil {
		logger.Error(err.Error())
		return nil, ErrGetConnection
	}

	decryptedConnections, err := crypto.Decrypt(encryptedConnections, cryptKey)
	if err != nil {
		logger.Error(err.Error())
		return nil, ErrDecryptData
	}

	err = json.Unmarshal([]byte(decryptedConnections), &connections)
	if err != nil {
		logger.Error(err.Error())
		return nil, ErrUnmarshalJson
	}

	return &connections, nil
}

// SetConnections save connection in file
func SetConnections(connections *connect.Connections) error {
	cryptKey, err := GetCryptKey()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	jsonConnections, err := json.Marshal(connections)
	if err != nil {
		logger.Error(err.Error())
		return ErrMarshalJson
	}

	updatedEncryptedConnections, err := crypto.Encrypt(string(jsonConnections), cryptKey)
	if err != nil {
		logger.Error(err.Error())
		return ErrEncryptData
	}

	err = storage.Write(DirectionApp, FileConnections, updatedEncryptedConnections)
	if err != nil {
		logger.Error(err.Error())
		return ErrWriteJson
	}

	return nil
}
