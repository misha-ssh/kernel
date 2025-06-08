package fuzz

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/crypto"
	"github.com/stretchr/testify/assert"
)

func FuzzCrypto_Encrypt(f *testing.F) {
	key, _ := crypto.GenerateKey()

	f.Fuzz(func(t *testing.T, ciphertext string) {
		_, err := crypto.Encrypt(ciphertext, key)
		assert.NoError(t, err)
	})
}
