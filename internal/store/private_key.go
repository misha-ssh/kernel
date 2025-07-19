package store

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"reflect"

	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"golang.org/x/crypto/ssh"
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

	if _, err := ssh.ParseRawPrivateKey([]byte(privateKey)); err == nil {
		return nil
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

// SavePrivateKey create private key for connection in spec dir
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

// UpdatePrivateKey update data private key
func UpdatePrivateKey(connection *connect.Connect) (string, error) {
	existFilenamePrivateKey := connection.Alias

	if !storage.Exists(DirectionKeys, existFilenamePrivateKey) {
		if len(connection.SshOptions.PrivateKey) == 0 {
			return "", nil
		}

		return SavePrivateKey(connection)
	}

	if len(connection.SshOptions.PrivateKey) == 0 {
		err := DeletePrivateKey(connection)
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}

		return "", nil
	}

	existDataPrivateKey, err := storage.Get(DirectionKeys, existFilenamePrivateKey)
	if err != nil {
		logger.Error(ErrGetDataPrivateKey.Error())
		return "", ErrGetDataPrivateKey
	}

	direction, filename := storage.GetDirectionAndFilename(connection.SshOptions.PrivateKey)
	dataPrivateKey, err := storage.Get(direction, filename)
	if err != nil {
		logger.Error(ErrGetDataPrivateKey.Error())
		return "", ErrGetDataPrivateKey
	}

	if !reflect.DeepEqual(existDataPrivateKey, dataPrivateKey) {
		err = DeletePrivateKey(connection)
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}

		return SavePrivateKey(connection)
	}

	return storage.GetFullPath(DirectionKeys, existFilenamePrivateKey), nil
}

// DeletePrivateKey delete key from dir for current connection
func DeletePrivateKey(connection *connect.Connect) error {
	directionPrivateKey := storage.GetPrivateKeysDir()
	filenamePrivateKey := connection.Alias

	return storage.Delete(directionPrivateKey, filenamePrivateKey)
}
