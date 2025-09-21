package testutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
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

	file := filepath.Join(direction, "key")
	err = os.WriteFile(file, privatePEM, 0644)
	if err != nil {
		return "", err
	}

	return file, nil
}

// CreateInvalidPrivateKey create invalid key for tests
func CreateInvalidPrivateKey(direction string) (string, error) {
	file := filepath.Join(direction, "invalid")
	err := os.WriteFile(file, []byte(""), 0644)
	if err != nil {
		return "", err
	}

	return file, nil
}
