package logger

import (
	"errors"
	"math/rand"
	"os"
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestStorageLogger_Error(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(*storage.MockStorage)
		isError   bool
		value     any
	}{
		{
			name: "success logging function",
			setupMock: func(m *storage.MockStorage) {
				m.On("Create", NameLogFile).Return(nil)

				file, _ := os.CreateTemp("", "log.log")
				m.On("GetOpenFile", NameLogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE).Return(file, nil)
			},
			isError: false,
			value:   rand.Int(),
		},
		{
			name: "error on file creation",
			setupMock: func(m *storage.MockStorage) {
				m.On("Create", NameLogFile).Return(errors.New("create error"))
			},
			isError: true,
			value:   rand.Int(),
		},
		{
			name: "error on getting open file",
			setupMock: func(m *storage.MockStorage) {
				m.On("Create", NameLogFile).Return(nil)
				m.On("GetOpenFile", NameLogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE).Return(nil, errors.New("open error"))
			},
			isError: true,
			value:   rand.Int(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := new(storage.MockStorage)
			tt.setupMock(mockStorage)

			sl := &StorageLogger{
				Storage: mockStorage,
			}

			if tt.isError {
				assert.Panics(t, func() {
					sl.Error(tt.value)
				})
			} else {
				sl.Error(tt.value)
			}

			mockStorage.AssertExpectations(t)
		})
	}
}

func TestStorageLogger_log(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(*storage.MockStorage)
		value     any
		wantErr   bool
	}{
		{
			name: "success logging function",
			setupMock: func(m *storage.MockStorage) {
				m.On("Create", NameLogFile).Return(nil)

				file, _ := os.CreateTemp("", "log.log")
				defer func() {
					err := os.Remove(file.Name())
					assert.NoError(t, err)
				}()

				m.On("GetOpenFile", NameLogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE).Return(file, nil)
			},
			value:   rand.Int(),
			wantErr: false,
		},
		{
			name: "error on file creation",
			setupMock: func(m *storage.MockStorage) {
				m.On("Create", NameLogFile).Return(errors.New("create error"))
			},
			value:   rand.Int(),
			wantErr: true,
		},
		{
			name: "error on getting open file",
			setupMock: func(m *storage.MockStorage) {
				m.On("Create", NameLogFile).Return(nil)
				m.On("GetOpenFile", NameLogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE).Return(nil, errors.New("open error"))
			},
			value:   rand.Int(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := new(storage.MockStorage)
			tt.setupMock(mockStorage)

			sl := &StorageLogger{
				Storage: mockStorage,
			}

			err := sl.log(tt.value)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockStorage.AssertExpectations(t)
		})
	}
}
