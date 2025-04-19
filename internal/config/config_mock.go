package config

import "github.com/stretchr/testify/mock"

type MockConfig struct {
	mock.Mock
}

func (m *MockConfig) Get(key string) string {
	args := m.Called(key)
	return args.String(0)
}

func (m *MockConfig) Set(key, value string) error {
	args := m.Called(key, value)
	return args.Error(0)
}
func (m *MockConfig) Exists(key string) bool {
	args := m.Called(key)
	return args.Bool(0)
}
