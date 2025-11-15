package space

import (
	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/configs/envname"
	"github.com/misha-ssh/kernel/internal/config"
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/ssh"
)

type Space struct {
	Storage   storage.Storage
	SSHConfig *ssh.Config
}

func New() *Space {
	space := new(Space)

	switch config.Get(envname.Space) {
	case envconst.TypeLocalSpace:
		space.Storage = storage.NewLocal()
	case envconst.TypeSSHConfig:
		space.SSHConfig = ssh.NewConfig()
	default:
		panic(envname.Space + ": unknown variable type or undefined")
	}

	return space
}

func (s *Space) GetConnections() (*connect.Connections, error) {
	return nil, nil
}

func (s *Space) SaveConnection(connection *connect.Connect) error {
	return nil
}

func (s *Space) UpdateConnection(connection *connect.Connect) (*connect.Connect, error) {
	return nil, nil
}

func (s *Space) DeleteConnection(connection *connect.Connect) error {
	return nil
}
