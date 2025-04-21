package config

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func FuzzLocalStorage_Set(f *testing.F) {
	localStorage := storage.LocalStorage{
		Direction: f.TempDir(),
	}

	f.Fuzz(func(t *testing.T, value string) {
		key := "TEST"

		s := StorageConfig{
			Storage: &localStorage,
		}

		_ = s.Set(key, value)

		got := s.Get(key)
		validateErr := validateValue(value)

		if validateErr != nil {
			assert.Equal(t, EmptyValue, got)
		} else {
			assert.Equal(t, value, got)
		}
	})
}
