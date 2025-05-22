package crypto

type Crypto interface {
	Encrypt(plaintext string, key string) (string, error)
	Decrypt(ciphertext string, key string) (string, error)
	GenerateKey() (string, error)
}
