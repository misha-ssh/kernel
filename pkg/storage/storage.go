package storage

import (
	"os"

	"github.com/stretchr/testify/mock"
)

type Storage interface {
	Exists(filename string, direction string) bool
	Create(filename string, direction string) error
	Get(filename string, direction string) (string, error)
	Delete(filename string, direction string) error
	Write(filename string, direction string, data string) error
	GetOpenFile(filename string, direction string) (*os.File, error)
}

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) Exists(filename string, direction string) bool {
	args := m.Called(filename, direction)
	return args.Bool(0)
}

func (m *MockStorage) Create(filename string, direction string) error {
	args := m.Called(filename, direction)
	return args.Error(0)
}

func (m *MockStorage) Get(filename string, direction string) (string, error) {
	args := m.Called(filename, direction)
	return args.String(0), args.Error(1)
}

func (m *MockStorage) Delete(filename string, direction string) error {
	args := m.Called(filename, direction)
	return args.Error(0)
}

func (m *MockStorage) Write(filename string, direction string, data string) error {
	args := m.Called(filename, direction, data)
	return args.Error(0)
}

func (m *MockStorage) GetOpenFile(filename string, direction string) (*os.File, error) {
	args := m.Called(filename, direction)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*os.File), args.Error(1)
}
