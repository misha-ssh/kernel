package fuzz

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"github.com/stretchr/testify/assert"
)

func FuzzStorage_Write(f *testing.F) {
	f.Fuzz(func(t *testing.T, value string) {
		direction := t.TempDir()
		fileName := "test"

		err := storage.Write(direction, fileName, value)
		assert.NoError(t, err)

		got, err := storage.Get(direction, fileName)
		assert.Equal(t, value, got)
	})
}
