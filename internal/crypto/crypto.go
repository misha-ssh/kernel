package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"log"

	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
)

const (
	SizeKey      = 32
	FilenameSalt = "salt.txt"
)

var (
	ErrVerifyKEy      = errors.New("err verify key on standard aes")
	ErrBlockCipher    = errors.New("err create 128-bit block cipher")
	ErrRandRead       = errors.New("err rand read encrypt")
	ErrAuthCiphertext = errors.New("err open decrypts and authenticates ciphertext")
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
		logger.Error(err)
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		logger.Error(err)
		return "", ErrRandRead
	}

	encryptData := string(gcm.Seal(nonce, nonce, []byte(plaintext), nil))

	return encryptData, nil
}

func Decrypt(ciphertext string, key string) (string, error) {
	ciphertextToByte := []byte(ciphertext)

	gcm, err := getGcm(key)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertextToByte := ciphertextToByte[:nonceSize], ciphertextToByte[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertextToByte, nil)
	if err != nil {
		logger.Error(err)
		return "", ErrAuthCiphertext
	}

	return string(plaintext), nil
}

// todo здесь реализовать генерациб соли и сохранения ее файл - если есть то отдача соли из файла
// так же можно добавить 600 права на файл
func getSalt() {

}

// GetKey todo написать генерацию соли + сохранения в файле - чтобы можно было получить ключ из пароля
func GetKey(password string) (string, error) {
	salt := make([]byte, 16)

	if _, err := rand.Read(salt); err != nil {
		log.Fatal(err)
	}

	return string(""), nil
}
