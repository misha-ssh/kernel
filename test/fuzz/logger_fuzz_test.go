package fuzz

import (
	"testing"

	"github.com/misha-ssh/kernel/internal/logger"
)

func FuzzLogger_Error(f *testing.F) {
	f.Fuzz(func(t *testing.T, value []uint8) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Error() is panicked, value = %v", value)
			}
		}()

		logger.Error(value)
	})
}
