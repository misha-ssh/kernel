package config

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func FuzzLocalStorage_Set(f *testing.F) {
	f.Fuzz(func(t *testing.T, value string) {
		localStorage := storage.LocalStorage{
			Direction: t.TempDir(),
		}

		s := StorageConfig{
			Storage: &localStorage,
		}

		key := "test"

		err := s.Set(key, value)
		got := s.Get(key)

		if err == nil {
			assert.Equal(t, value, got)
		}
	})
}
