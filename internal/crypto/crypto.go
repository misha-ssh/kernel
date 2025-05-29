package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"golang.org/x/crypto/scrypt"
)

const (
	SizeKey         = 32
	Cost            = 32768
	BlockSize       = 8
	Parallelization = 1

	FilenameSalt = "salt"
)

var (
	ErrVerifyKEy          = errors.New("err verify key on standard aes")
	ErrBlockCipher        = errors.New("err create 128-bit block cipher")
	ErrRandRead           = errors.New("err rand read encrypt")
	ErrAuthCiphertext     = errors.New("err open decrypts and authenticates ciphertext")
	ErrGenerateRandomSalt = errors.New("err generate random salt")
	ErrCreateSaltFile     = errors.New("err create salt file")
	ErrWriteInFileSalt    = errors.New("err write in salt file")
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

func getSalt() (string, error) {
	homeStorage := storage.FileStorage{
		Direction: storage.GetHomeDir(),
	}

	if !homeStorage.Exists(FilenameSalt) {
		err := homeStorage.Create(FilenameSalt)
		if err != nil {
			logger.Error(ErrCreateSaltFile)
			return "", ErrCreateSaltFile
		}

		salt := make([]byte, 16)

		if _, err := rand.Read(salt); err != nil {
			logger.Error(ErrGenerateRandomSalt)
			return "", ErrGenerateRandomSalt
		}

		err = homeStorage.Write(FilenameSalt, string(salt))
		if err != nil {
			logger.Error(ErrWriteInFileSalt)
			return "", ErrWriteInFileSalt
		}
	}

	salt, err := homeStorage.Get(FilenameSalt)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return salt, nil
}

func GetKey(password string) (string, error) {
	salt, err := getSalt()
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	key, err := scrypt.Key([]byte(password), []byte(salt), Cost, BlockSize, Parallelization, SizeKey)

	return string(key), nil
}
