package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

const SizeKey = 64

var (
	ErrGenerateKey = errors.New("err at created log file")
)

type StorageEncryption struct{}

func (s *StorageEncryption) Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	newCipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(newCipher)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func (s *StorageEncryption) Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	newCipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(newCipher)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	return plaintext, nil
}

func (s *StorageEncryption) GenerateKey() ([]byte, error) {
	key := make([]byte, SizeKey)

	_, err := rand.Read(key)
	if err != nil {
		return nil, ErrGenerateKey
	}

	return key, nil
}
