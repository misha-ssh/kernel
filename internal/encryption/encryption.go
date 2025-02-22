package encryption

type Encryption interface {
	Encrypt(plaintext []byte, key []byte) ([]byte, error)
	Decrypt(ciphertext []byte, key []byte) ([]byte, error)
	GenerateKey() ([]byte, error)
	GetKey() ([]byte, error)
}
