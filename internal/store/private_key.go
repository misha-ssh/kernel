package store

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

var (
	DirectionKeys = storage.GetPrivateKeysDir()

	ErrWriteToFilePrivateKey = errors.New("err write to file private key")
	ErrCreateFilePrivateKey  = errors.New("err create file private key")
	ErrNotValidPrivateKey    = errors.New("private key is not valid")
	ErrGetDataPrivateKey     = errors.New("private key get data error")
)

func validatePrivateKey(privateKey string) error {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return ErrNotValidPrivateKey
	}

	_, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return ErrNotValidPrivateKey
		}

		_, ok := key.(*rsa.PrivateKey)
		if !ok {
			return ErrNotValidPrivateKey
		}

		return nil
	}

	return nil
}

func SavePrivateKey(connection *connect.Connect) (string, error) {
	direction, filename := storage.GetDirectionAndFilename(connection.SshOptions.PrivateKey)
	dataPrivateKey, err := storage.Get(direction, filename)
	if err != nil {
		logger.Error(ErrGetDataPrivateKey.Error())
		return "", ErrGetDataPrivateKey
	}

	err = validatePrivateKey(dataPrivateKey)
	if err != nil {
		logger.Error(ErrNotValidPrivateKey.Error())
		return "", ErrNotValidPrivateKey
	}

	filenamePrivateKey := connection.Alias

	err = storage.Create(DirectionKeys, filenamePrivateKey)
	if err != nil {
		logger.Error(ErrCreateFilePrivateKey.Error())
		return "", ErrCreateFilePrivateKey
	}

	err = storage.Write(DirectionKeys, filenamePrivateKey, dataPrivateKey)
	if err != nil {
		logger.Error(ErrWriteToFilePrivateKey.Error())
		return "", ErrWriteToFilePrivateKey
	}

	return storage.GetFullPath(DirectionKeys, filenamePrivateKey), nil
}

func DeletePrivateKey(connection *connect.Connect) error {
	directionPrivateKey := storage.GetPrivateKeysDir()
	filenamePrivateKey := connection.Alias

	return storage.Delete(directionPrivateKey, filenamePrivateKey)
}
