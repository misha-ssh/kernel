package crypt

import "errors"

func Encrypt(str string) (string, error) {
	cryptKey, err := GetKey()
	if err != nil {
		return "", errors.New("err at get crypt key")
	}

	ciphertext, err := encrypt(cryptKey, str)
	if err != nil {
		return "", errors.New("encryption error")
	}

	return ciphertext, nil
}
