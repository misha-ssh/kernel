package testutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

func CreateInvalidPrivateKey(direction string) (string, error) {
	filenameInvalidKey := "invalid"
	err := storage.Write(direction, filenameInvalidKey, "")
	if err != nil {
		return "", err
	}

	return storage.GetFullPath(direction, filenameInvalidKey), nil
}

func CreatePrivateKey(direction string) (string, error) {
	filenameKey := "key"

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", err
	}

	privateDER := x509.MarshalPKCS1PrivateKey(privateKey)

	privateBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateDER,
	}

	privatePEM := pem.EncodeToMemory(&privateBlock)

	err = storage.Write(direction, filenameKey, string(privatePEM))
	if err != nil {
		return "", err
	}

	return storage.GetFullPath(direction, filenameKey), nil
}
