package crypt

import (
	"crypto/rand"
	"errors"

	"github.com/ssh-connection-manager/file"
)

func GetKey() ([]byte, error) {
	fileCrypt := GetFile()

	data, err := fileCrypt.ReadFile()
	if err != nil {
		return []byte(data), errors.New("empty name")
	}

	return []byte(data), nil
}

func GenerateFileKey(fl file.File) error {
	SetFile(fl)
	fileKey := GetFile()

	if !fileKey.IsExistFile() {
		err := fileKey.CreateFile()
		if err != nil {
			return err
		}

		cryptKey, err := GetKey()
		if err != nil {
			return errors.New("empty name")
		}

		if len(cryptKey) == 0 {
			keyData := make([]byte, 32)

			_, err := rand.Read(keyData)
			if err != nil {
				return errors.New("key generation error")
			}

			err = fileKey.WriteFile(keyData)
			if err != nil {
				return errors.New("error writing key")
			}
		}
	}

	return nil
}
