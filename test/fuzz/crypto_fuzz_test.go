package fuzz

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/internal/crypto"
)

func FuzzCrypto_Encrypt(f *testing.F) {
	key, _ := crypto.GenerateKey()

	f.Fuzz(func(t *testing.T, ciphertext string) {
		_, err := crypto.Encrypt(ciphertext, key)
		if err != nil {
			t.Errorf("encrypt failed: %v", err)
		}
	})
}
