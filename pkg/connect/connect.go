package connect

type Connections struct {
	Connects []Connect `json:"connects"`
}

// Connect represents a single connection configuration
type Connect struct {
	// Alias is a user-defined name for the connection
	Alias string `json:"alias"`
	// Login is the username for authentication
	Login string `json:"login"`
	// Address is the hostname or IP address of the remote server
	Address string `json:"address"`
	// Password is the password for authentication
	Password string `json:"password"`

	// CreatedAt is the timestamp when this connection was created
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when this connection was last modified
	UpdatedAt string `json:"updated_at"`

	// Port specifies the SSH port
	Port int `json:"port"`
	// PrivateKey contains the PEM-encoded private key for authentication
	PrivateKey string `json:"private_key"`
	// Passphrase is the passphrase for decrypting the private key
	Passphrase string `json:"passphrase"`
}
