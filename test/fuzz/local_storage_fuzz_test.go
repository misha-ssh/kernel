package fuzz

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func FuzzLocalStorage_Write(f *testing.F) {
	f.Fuzz(func(t *testing.T, value string) {
		s := storage.LocalStorage{
			Direction: t.TempDir(),
		}

		fileName := "test"

		err := s.Write(fileName, value)
		assert.NoError(t, err)

		got, err := s.Get(fileName)
		assert.Equal(t, value, got)
	})
}
