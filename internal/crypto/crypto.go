package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"

	"github.com/ssh-connection-manager/kernel/internal/logger"
)

const (
	SizeKey = 32
)

var (
	ErrVerifyKEy      = errors.New("err verify key on standard aes")
	ErrBlockCipher    = errors.New("err create 128-bit block cipher")
	ErrRandRead       = errors.New("err rand read encrypt")
	ErrAuthCiphertext = errors.New("err open decrypts and authenticates ciphertext")
	ErrGenerateKey    = errors.New("err generate key")
)

func getGcm(key string) (cipher.AEAD, error) {
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

// Encrypt securely encrypts plaintext using AES-GCM with:
// - Provided encryption key
// - Random nonce generation
// Returns ciphertext or error if encryption fails
func Encrypt(plaintext string, key string) (string, error) {
	gcm, err := getGcm(key)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		logger.Error(err.Error())
		return "", ErrRandRead
	}

	encryptData := string(gcm.Seal(nonce, nonce, []byte(plaintext), nil))

	return encryptData, nil
}

// Decrypt authenticates and decrypts ciphertext using:
// - Same key used for encryption
// - Embedded nonce from ciphertext
// Returns plaintext or error if decryption/authentication fails
func Decrypt(ciphertext string, key string) (string, error) {
	ciphertextToByte := []byte(ciphertext)

	gcm, err := getGcm(key)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertextToByte := ciphertextToByte[:nonceSize], ciphertextToByte[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertextToByte, nil)
	if err != nil {
		logger.Error(err.Error())
		return "", ErrAuthCiphertext
	}

	return string(plaintext), nil
}

// GenerateKey creates a cryptographically secure:
// - 256-bit (32-byte) random key
// - Suitable for AES-256 encryption
// Returns key or error if random generation fails
func GenerateKey() (string, error) {
	key := make([]byte, SizeKey)

	_, err := rand.Read(key)
	if err != nil {
		logger.Error(err.Error())
		return "", ErrGenerateKey
	}

	return string(key), nil
}
