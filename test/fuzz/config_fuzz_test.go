package fuzz

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/config"
	"github.com/stretchr/testify/assert"
)

func FuzzConfig_Set(f *testing.F) {
	f.Fuzz(func(t *testing.T, value string) {
		key := "test"

		err := config.Set(key, value)
		got := config.Get(key)

		if err == nil {
			assert.Equal(t, value, got)
		}
	})
}
