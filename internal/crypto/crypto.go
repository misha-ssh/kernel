package crypto

import (
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

type Crypto interface {
	Encrypt(plaintext string, key string) (string, error)
	Decrypt(ciphertext string, key string) (string, error)
	GenerateKey() (string, error)
	GetKey(storage storage.Storage) (string, error)
}
