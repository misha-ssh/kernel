package store

import (
	"errors"
	"os/user"

	"github.com/ssh-connection-manager/kernel/v2/configs/envconst"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/zalando/go-keyring"
)

var ErrGetCryptKey = errors.New("err get crypt key")

// GetCryptKey get master key from keyring
func GetCryptKey() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	username := currentUser.Username

	cryptKey, err := keyring.Get(envconst.NameServiceCryptKey, username)
	if err != nil {
		logger.Error(ErrGetCryptKey.Error())
		return "", err
	}

	return cryptKey, nil
}
