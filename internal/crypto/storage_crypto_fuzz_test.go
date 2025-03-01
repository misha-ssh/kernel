package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FuzzStorageCrypto_Encrypt(f *testing.F) {
	se := &StorageCrypto{}
	key, _ := se.GenerateKey()

	f.Fuzz(func(t *testing.T, ciphertext string) {
		_, err := se.Encrypt(ciphertext, key)
		assert.NoError(t, err)
	})
}
