package crypto

import (
	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"github.com/stretchr/testify/mock"
)

type MockCrypto struct {
	mock.Mock
}

func (m *MockCrypto) Encrypt(plaintext string, key string) (string, error) {
	args := m.Called(plaintext, key)

	if args.Error(1) != nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), args.Error(1)
}

func (m *MockCrypto) Decrypt(ciphertext string, key string) (string, error) {
	args := m.Called(ciphertext, key)
	if args.Error(1) != nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), args.Error(1)
}

func (m *MockCrypto) GenerateKey() (string, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), args.Error(1)
}

func (m *MockCrypto) GetKey(storage storage.Storage) (string, error) {
	args := m.Called(storage)
	if args.Error(1) != nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), args.Error(1)
}
