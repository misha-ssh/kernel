package crypt

import (
	"errors"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
)

func Decrypt(str string) (string, error) {
	cryptKey, err := GetKey()
	if err != nil {
		logger.Danger(err.Error())
		return "", errors.New("err at get crypt key")
	}

	decrypted, err := decrypt(cryptKey, str)
	if err != nil {
		logger.Danger(err.Error())
		return "", errors.New("encryption error")
	}

	return decrypted, nil
}
