package config

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestStorageConfig_Set(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value string
	}{
		{
			name:  "set default value",
			key:   "test",
			value: "1",
		},
		{
			name:  "set value with spaces",
			key:   "test 2",
			value: "2",
		},
		{
			name:  "set value with random register",
			key:   "RaRd",
			value: "2",
		},
	}

	configStorage := storage.LocalStorage{
		Direction: "/home/deniskorbakov/storage-config/",
	}

	s := StorageConfig{
		Storage: &configStorage,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.Set(tt.key, tt.value)
			got := s.Get(tt.key)

			assert.Equal(t, tt.value, got)
		})
	}
}
