package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FuzzLocalStorage_Write(f *testing.F) {
	f.Fuzz(func(t *testing.T, value string) {
		s := LocalStorage{
			direction: t.TempDir(),
		}

		fileName := "test"

		err := s.Write(fileName, value)
		assert.NoError(t, err)

		got, err := s.Get(fileName)
		assert.Equal(t, value, got)
	})
}
