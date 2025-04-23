package fuzz

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/crypto"
	"github.com/stretchr/testify/assert"
)

func FuzzStorageCrypto_Encrypt(f *testing.F) {
	se := &crypto.StorageCrypto{}
	key, _ := se.GenerateKey()

	f.Fuzz(func(t *testing.T, ciphertext string) {
		_, err := se.Encrypt(ciphertext, key)
		assert.NoError(t, err)
	})
}
