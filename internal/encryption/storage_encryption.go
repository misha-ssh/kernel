package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

const SizeKey = 64

var (
	ErrGenerateKey    = errors.New("err at created log file")
	ErrVerifyKEy      = errors.New("err verify key on standard aes")
	ErrBlockCipher    = errors.New("err create 128-bit block cipher")
	ErrRandRead       = errors.New("err rand read encrypt")
	ErrAuthCiphertext = errors.New("err open decrypts and authenticates ciphertext")
)

type StorageEncryption struct{}

func (s *StorageEncryption) getGcm(key []byte) (cipher.AEAD, error) {
	newCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, ErrVerifyKEy
	}

	gcm, err := cipher.NewGCM(newCipher)
	if err != nil {
		return nil, ErrBlockCipher
	}

	return gcm, nil
}

func (s *StorageEncryption) Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	gcm, err := s.getGcm(key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, ErrRandRead
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func (s *StorageEncryption) Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	gcm, err := s.getGcm(key)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, ErrAuthCiphertext
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

func (s *StorageEncryption) GetKey() ([]byte, error) {
	return nil, nil
}
