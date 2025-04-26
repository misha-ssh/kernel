package fuzz

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/stretchr/testify/assert"
)

func FuzzStorageLogger_Error(f *testing.F) {
	f.Fuzz(func(t *testing.T, value []uint8) {
		assert.NotPanics(t, func() {
			logger.e(value)
		})
	})
}
