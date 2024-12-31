package crypt

import (
	"crypto/rand"
	"errors"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"

	"github.com/ssh-connection-manager/kernel/v2/pkg/file"
)

func GetKey() ([]byte, error) {
	fileCrypt := GetFile()

	data, err := fileCrypt.ReadFile()
	if err != nil {
		errText := "empty name"

		logger.Danger(errText)
		return []byte(data), errors.New(errText)
	}

	return []byte(data), nil
}

func GenerateFileKey(fl file.File) error {
	SetFile(fl)
	fileKey := GetFile()

	if !fileKey.IsExistFile() {
		err := fileKey.CreateFile()
		if err != nil {
			logger.Danger(err.Error())
			return err
		}

		cryptKey, err := GetKey()
		if err != nil {
			logger.Danger(err.Error())
			return errors.New("empty name")
		}

		if len(cryptKey) == 0 {
			keyData := make([]byte, 32)

			_, err := rand.Read(keyData)
			if err != nil {
				logger.Danger(err.Error())
				return errors.New("key generation error")
			}

			err = fileKey.WriteFile(keyData)
			if err != nil {
				logger.Danger(err.Error())
				return errors.New("error writing key")
			}
		}
	}

	return nil
}
