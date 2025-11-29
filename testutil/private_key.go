package testutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
)

// generatePrivateKey generate private key for ssh connect
func generatePrivateKey(pass string) ([]byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}

	privateBlock := &pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(privateKey),
	}

	if pass != "" {
		//nolint:all
		privateBlock, err = x509.EncryptPEMBlock(rand.Reader, privateBlock.Type, privateBlock.Bytes, []byte(pass), x509.PEMCipherAES256)
		if err != nil {
			return nil, err
		}
	}

	privatePEM := pem.EncodeToMemory(privateBlock)

	return privatePEM, nil
}

// CreatePrivateKey generate and save private key in file
func CreatePrivateKey(direction string) (string, error) {
	privatePEM, err := generatePrivateKey("")
	if err != nil {
		return "", err
	}

	file := filepath.Join(direction, "key")
	err = os.WriteFile(file, privatePEM, os.ModePerm)
	if err != nil {
		return "", err
	}

	return file, nil
}

// CreatePrivateKeyWithPass generate and save private key with pass in file
func CreatePrivateKeyWithPass(direction string, pass string) (string, error) {
	privatePEM, err := generatePrivateKey(pass)
	if err != nil {
		return "", err
	}

	file := filepath.Join(direction, "key-pass")
	err = os.WriteFile(file, privatePEM, os.ModePerm)
	if err != nil {
		return "", err
	}

	return file, nil
}

// CreateInvalidPrivateKey create invalid key for tests
func CreateInvalidPrivateKey(direction string) (string, error) {
	file := filepath.Join(direction, "invalid")
	err := os.WriteFile(file, []byte(""), os.ModePerm)
	if err != nil {
		return "", err
	}

	return file, nil
}
