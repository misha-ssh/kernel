package encryption

import (
	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"github.com/stretchr/testify/mock"
)

type Encryption interface {
	Encrypt(plaintext string, key string) (string, error)
	Decrypt(ciphertext string, key string) (string, error)
	GenerateKey() (string, error)
	GetKey(storage storage.Storage) (string, error)
}

type MockEncryption struct {
	mock.Mock
}

func (m *MockEncryption) Encrypt(plaintext string, key string) (string, error) {
	args := m.Called(plaintext, key)

	if args.Error(1) != nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), args.Error(1)
}

func (m *MockEncryption) Decrypt(ciphertext string, key string) (string, error) {
	args := m.Called(ciphertext, key)
	if args.Error(1) != nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), args.Error(1)
}

func (m *MockEncryption) GenerateKey() (string, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), args.Error(1)
}

func (m *MockEncryption) GetKey(storage storage.Storage) (string, error) {
	args := m.Called(storage)
	if args.Error(1) != nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), args.Error(1)
}
