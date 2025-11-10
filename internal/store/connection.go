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
const FileConfigSSH = envconst.FilenameConfigSSH

var (
	DirectionApp = storage.GetAppDir()
	DirectionSSH = storage.GetDirSSH()

	ErrEncryptData   = errors.New("err encrypt data")
	ErrMarshalJson   = errors.New("failed to marshal json")
	ErrWriteJson     = errors.New("failed to write json")
	ErrUnmarshalJson = errors.New("failed to unmarshal json")
	ErrGetConnection = errors.New("failed to get connection")
	ErrDecryptData   = errors.New("failed to decrypt data")
)

func getFromLocalStorage() (string, error) {
	cryptKey, err := GetCryptKey()
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	encryptedConnections, err := storage.Get(DirectionApp, FileConnections)
	if err != nil {
		logger.Error(err.Error())
		return "", ErrGetConnection
	}

	decryptedConnections, err := crypto.Decrypt(encryptedConnections, cryptKey)
	if err != nil {
		logger.Error(err.Error())
		return "", ErrDecryptData
	}

	return decryptedConnections, nil
}

// todo add logic get connection from ssh config
func getFromConfigSSH() (string, error) {
	sshConfig, err := storage.Get(DirectionSSH, FileConfigSSH)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return sshConfig, nil
}

// GetConnections get connection from file
func GetConnections() (*connect.Connections, error) {
	var connections connect.Connections

	//todo add fast read ~goroutine

	cls, err := getFromLocalStorage()
	if err != nil {
		logger.Error(err.Error())
		return nil, ErrGetConnection
	}

	err = json.Unmarshal([]byte(cls), &connections)
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
