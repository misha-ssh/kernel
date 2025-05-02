package config

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestStorage_Set(t *testing.T) {
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
			name:    "set char early",
			key:     "testEarly",
			value:   "=",
			wantErr: true,
		},
		{
			name:    "set char early - rewrite key",
			key:     "test",
			value:   "=",
			wantErr: true,
		},
		{
			name:    "set numeric",
			key:     "randomKey",
			value:   "0",
			wantErr: false,
		},
		{
			name:    "set one char",
			key:     "test",
			value:   "z",
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
			value:   "test",
			wantErr: false,
		},
	}

	configStorage := storage.FileStorage{
		Direction: t.TempDir(),
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

func TestStorage_Get(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		want       string
		isSetValue bool
	}{
		{
			name:       "get existing value",
			key:        "test",
			want:       "testData",
			isSetValue: true,
		},
		{
			name:       "get rewriting value",
			key:        "test",
			want:       "rewriteData",
			isSetValue: true,
		},
		{
			name:       "get numeric value",
			key:        "newTestKey",
			want:       "0",
			isSetValue: true,
		},
		{
			name:       "get lat set value",
			key:        "test",
			want:       "rewriteData",
			isSetValue: false,
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

	configStorage := storage.FileStorage{
		Direction: t.TempDir(),
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

func TestStorage_Exists(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		isCreateKey bool
		want        bool
	}{
		{
			name:        "key exists",
			key:         "test",
			isCreateKey: true,
			want:        true,
		},
		{
			name:        "key exists with random register",
			key:         "tEsT",
			isCreateKey: false,
			want:        true,
		},
		{
			name:        "key dont exists",
			key:         "dontExistKey",
			isCreateKey: false,
			want:        false,
		},
		{
			name:        "empty key",
			key:         "",
			isCreateKey: false,
			want:        false,
		},
	}

	configStorage := storage.FileStorage{
		Direction: t.TempDir(),
	}

	s := StorageConfig{
		Storage: &configStorage,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isCreateKey {
				err := s.Set(tt.key, "test")
				assert.NoError(t, err)
			}

			got := s.Exists(tt.key)
			assert.Equal(t, tt.want, got)
		})
	}
}
