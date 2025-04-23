package fuzz

import (
	"os"
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func FuzzStorageLogger_LocStorageErr(f *testing.F) {
	mockStorage := new(storage.MockStorage)
	mockStorage.On("Create", logger.NameLogFile).Return(nil)

	file, _ := os.CreateTemp("", "log.log")
	defer func() {
		err := os.Remove(file.Name())
		assert.NoError(f, err)
	}()

	mockStorage.On("GetOpenFile", logger.NameLogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE).Return(file, nil)

	sl := &logger.StorageLogger{
		Storage: mockStorage,
	}

	f.Fuzz(func(t *testing.T, value []uint8) {
		assert.NotPanics(t, func() {
			sl.Error(value)
		})
	})
}
