package storage

import (
	"github.com/stretchr/testify/mock"
	"os"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) Exists(filename string) bool {
	args := m.Called(filename)
	return args.Bool(0)
}

func (m *MockStorage) Create(filename string) error {
	args := m.Called(filename)
	return args.Error(0)
}

func (m *MockStorage) Get(filename string) (string, error) {
	args := m.Called(filename)
	return args.String(0), args.Error(1)
}

func (m *MockStorage) Delete(filename string) error {
	args := m.Called(filename)
	return args.Error(0)
}

func (m *MockStorage) Write(filename string, data string) error {
	args := m.Called(filename, data)
	return args.Error(0)
}

func (m *MockStorage) GetOpenFile(filename string) (*os.File, error) {
	args := m.Called(filename)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*os.File), args.Error(1)
}
