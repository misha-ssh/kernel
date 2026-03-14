package store

import (
	"errors"
	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"reflect"
)

var (
	DirectionKeys = storage.GetPrivateKeysDir()

	ErrWriteToFilePrivateKey = errors.New("err write to file private key")
	ErrCreateFilePrivateKey  = errors.New("err create file private key")
	ErrGetDataPrivateKey     = errors.New("private key get data error")
)

// SavePrivateKey create private key for connection in spec dir
func SavePrivateKey(connection *connect.Connect) (string, error) {
	direction, filename := storage.GetDirectionAndFilename(connection.SshOptions.PrivateKey)
	dataPrivateKey, err := storage.Get(direction, filename)
	if err != nil {
		logger.Error(err.Error())
		return "", ErrGetDataPrivateKey
	}

	filenamePrivateKey := connection.Alias

	err = storage.Create(DirectionKeys, filenamePrivateKey)
	if err != nil {
		logger.Error(err.Error())
		return "", ErrCreateFilePrivateKey
	}

	err = storage.Write(DirectionKeys, filenamePrivateKey, dataPrivateKey)
	if err != nil {
		logger.Error(err.Error())
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
		logger.Error(err.Error())
		return "", ErrGetDataPrivateKey
	}

	direction, filename := storage.GetDirectionAndFilename(connection.SshOptions.PrivateKey)
	dataPrivateKey, err := storage.Get(direction, filename)
	if err != nil {
		logger.Error(err.Error())
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
