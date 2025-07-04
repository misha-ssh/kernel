package fuzz

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

func FuzzStorage_Write(f *testing.F) {
	f.Fuzz(func(t *testing.T, value string) {
		direction := t.TempDir()
		fileName := "test"

		err := storage.Write(direction, fileName, value)
		if err != nil {
			t.Errorf("write failed: %v", err)
		}

		got, err := storage.Get(direction, fileName)
		if err != nil {
			t.Errorf("Get failed: %v", err)
		}

		if got != value {
			t.Errorf("got %q != want %q", got, value)
		}
	})
}
