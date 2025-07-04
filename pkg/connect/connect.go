package connect

import "golang.org/x/crypto/ssh"

type Connector interface {
	Connect(session *ssh.Session) error
	NewSession(connection *Connect) (*ssh.Session, error)
}

type ConnectionType string

// TypeSSH type for ssh connection
const TypeSSH ConnectionType = "ssh"

type Connections struct {
	Connects []Connect `json:"connects"`
}

// Connect represents a single connection configuration
type Connect struct {
	// Alias is a user-defined name for the connection
	Alias     string `json:"alias"`
	Login     string `json:"login"`
	Address   string `json:"address"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	// Type specifies the connection protocol (e.g., "ssh")
	Type ConnectionType `json:"type"`

	// SshOptions contains SSH-specific configuration options
	SshOptions *SshOptions `json:"ssh_options,omitempty"`
}

// SshOptions contains configuration options specific to SSH connections
type SshOptions struct {
	// Port specifies the SSH port (default is 22 if not specified)
	Port int `json:"port"`

	// PrivateKey contains the PEM-encoded private key for authentication
	PrivateKey string `json:"private_key"`
}
