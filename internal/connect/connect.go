package connect

type Connector interface {
	Connect(connection *Connect) error
}

type ConnectionType string

const (
	TypeSSH ConnectionType = "ssh"
	TypeFTP ConnectionType = "ftp"
)

type Connections struct {
	Connects []Connect `json:"connects"`
}

type Connect struct {
	Alias     string         `json:"alias"`
	Login     string         `json:"login"`
	Address   string         `json:"address"`
	Password  string         `json:"password"`
	Type      ConnectionType `json:"type"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`

	SshOptions *SshOptions `json:"ssh_options,omitempty"`
}

type SshOptions struct {
	Port       int  `json:"port"`
	PrivateKey bool `json:"private_key"`
}
