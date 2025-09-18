package testutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/misha-ssh/kernel/internal/storage"
)

// GeneratePrivateKey generate private key for ssh connect
func GeneratePrivateKey() ([]byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}

	privateDER := x509.MarshalPKCS1PrivateKey(privateKey)

	privateBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateDER,
	}

	privatePEM := pem.EncodeToMemory(&privateBlock)

	return privatePEM, nil
}

// CreatePrivateKey generate and save private key in file
func CreatePrivateKey(direction string) (string, error) {
	privatePEM, err := GeneratePrivateKey()
	if err != nil {
		return "", err
	}

	filenameKey := "key"

	err = storage.Write(direction, filenameKey, string(privatePEM))
	if err != nil {
		return "", err
	}

	return storage.GetFullPath(direction, filenameKey), nil
}

// CreateInvalidPrivateKey create invalid key for tests
func CreateInvalidPrivateKey(direction string) (string, error) {
	filenameInvalidKey := "invalid"
	err := storage.Write(direction, filenameInvalidKey, "")
	if err != nil {
		return "", err
	}

	return storage.GetFullPath(direction, filenameInvalidKey), nil
}
