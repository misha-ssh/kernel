package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
)

const (
	SizeKey  = 32
	FileName = "encryption.key"
)

var (
	ErrGenerateKey    = errors.New("err at created log file")
	ErrVerifyKEy      = errors.New("err verify key on standard aes")
	ErrBlockCipher    = errors.New("err create 128-bit block cipher")
	ErrRandRead       = errors.New("err rand read encrypt")
	ErrAuthCiphertext = errors.New("err open decrypts and authenticates ciphertext")
)

type StorageEncryption struct{}

func (s *StorageEncryption) getGcm(key string) (cipher.AEAD, error) {
	newCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, ErrVerifyKEy
	}

	gcm, err := cipher.NewGCM(newCipher)
	if err != nil {
		return nil, ErrBlockCipher
	}

	return gcm, nil
}

func (s *StorageEncryption) Encrypt(plaintext string, key string) (string, error) {
	gcm, err := s.getGcm(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", ErrRandRead
	}

	encryptData := string(gcm.Seal(nonce, nonce, []byte(plaintext), nil))

	return encryptData, nil
}

func (s *StorageEncryption) Decrypt(ciphertext string, key string) (string, error) {
	ciphertextToByte := []byte(ciphertext)

	gcm, err := s.getGcm(key)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertextToByte := ciphertextToByte[:nonceSize], ciphertextToByte[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertextToByte, nil)
	if err != nil {
		return "", ErrAuthCiphertext
	}

	return string(plaintext), nil
}

func (s *StorageEncryption) GenerateKey() (string, error) {
	key := make([]byte, SizeKey)

	_, err := rand.Read(key)
	if err != nil {
		return "", ErrGenerateKey
	}

	return string(key), nil
}

func (s *StorageEncryption) GetKey(storage storage.Storage) (string, error) {
	if storage.Exists(FileName) {
		key, err := storage.Get(FileName)
		if err != nil {
			return "", err
		}

		return key, nil
	}

	key, err := s.GenerateKey()
	if err != nil {
		return "", err
	}

	err = storage.Write(FileName, key)
	if err != nil {
		return "", err
	}

	return key, nil
}
