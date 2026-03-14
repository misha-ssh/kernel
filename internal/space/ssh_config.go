package space

import (
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/ssh"
)

type SSHConfig struct {
	Config *ssh.Config
}

func (s *SSHConfig) GetConnections() (*connect.Connections, error) {
	return nil, nil
}

func (s *SSHConfig) SaveConnection(connection *connect.Connect) error {
	return nil
}

func (s *SSHConfig) UpdateConnection(connection *connect.Connect) (*connect.Connect, error) {
	return nil, nil
}

func (s *SSHConfig) DeleteConnection(connection *connect.Connect) error {
	return nil
}
