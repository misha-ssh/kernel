package logger

import (
	"os"
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func FuzzStorageLogger_log(f *testing.F) {
	mockStorage := new(storage.MockStorage)
	mockStorage.On("Create", NameLogFile).Return(nil)

	file, _ := os.CreateTemp("", "log.log")
	defer func() {
		err := os.Remove(file.Name())
		assert.NoError(f, err)
	}()

	mockStorage.On("GetOpenFile", NameLogFile).Return(file, nil)

	sl := &StorageLogger{
		Storage: mockStorage,
	}

	f.Fuzz(func(t *testing.T, value []uint8) {
		err := sl.log(value)
		assert.NoError(t, err)
	})
}
