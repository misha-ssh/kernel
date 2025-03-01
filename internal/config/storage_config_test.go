package config

import (
	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"testing"
)

func TestStorageConfig_Set(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name      string
		setupMock func(*storage.MockStorage)
		args      args
	}{
		{
			name: "success",
			setupMock: func(m *storage.MockStorage) {
				m.On("Exists", FileName).Return(true)
				m.On("Write", "").Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := new(storage.MockStorage)
			tt.setupMock(mockStorage)

			l := &StorageConfig{
				Storage: mockStorage,
			}

			l.Set(tt.args.key, tt.args.value)
			mockStorage.AssertExpectations(t)
		})
	}
}
