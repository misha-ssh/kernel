package config

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestStorageConfig_Set(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		value   string
		wantErr bool
	}{
		{
			name:    "set numbers",
			key:     "test",
			value:   "123123",
			wantErr: false,
		},
		{
			name:    "set letters",
			key:     "test",
			value:   "randomString",
			wantErr: false,
		},
		{
			name:    "set letters and numbers",
			key:     "test",
			value:   "randomString123123",
			wantErr: false,
		},
		{
			name:    "empty key",
			key:     "",
			value:   "test",
			wantErr: true,
		},
		{
			name:    "key is spaces",
			key:     "   ",
			value:   "test",
			wantErr: true,
		},
		{
			name:    "empty value",
			key:     "test",
			value:   "",
			wantErr: true,
		},
		{
			name:    "values is spaces",
			key:     "test",
			value:   "   ",
			wantErr: true,
		},
		{
			name:    "set key with spaces",
			key:     "test 2",
			value:   "2",
			wantErr: true,
		},
		{
			name:    "set value with spaces",
			key:     "test2",
			value:   "a a",
			wantErr: true,
		},
		{
			name:    "set value with random register",
			key:     "RaRd",
			value:   "2",
			wantErr: false,
		},
	}

	//TODO: не забыть заменить на временную директорию
	configStorage := storage.LocalStorage{
		Direction: "/Users/deniskorbakov/storage-config-set/",
	}

	s := StorageConfig{
		Storage: &configStorage,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.Set(tt.key, tt.value)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				got := s.Get(tt.key)
				assert.Equal(t, tt.value, got)
			}
		})
	}
}

func TestStorageConfig_Get(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		want       string
		isSetValue bool
	}{
		{
			name:       "get existing value",
			key:        "test",
			want:       "testasdas",
			isSetValue: true,
		},
		{
			name:       "get non existing value",
			key:        "empty",
			want:       "",
			isSetValue: false,
		},
		{
			name:       "get empty key",
			key:        "",
			want:       "",
			isSetValue: false,
		},
	}

	//TODO: не забыть заменить на временную директорию
	configStorage := storage.LocalStorage{
		Direction: "/Users/deniskorbakov/storage-config-get/",
	}

	s := StorageConfig{
		Storage: &configStorage,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isSetValue {
				err := s.Set(tt.key, tt.want)
				assert.NoError(t, err)
			}

			got := s.Get(tt.key)
			assert.Equal(t, tt.want, got)
		})
	}
}
