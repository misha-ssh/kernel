package ssh

import (
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
)

type Config struct {
	Path string
}

func NewConfig() *Config {
	return &Config{
		Path: storage.GetDirSSH(),
	}
}

func (c *Config) GetConnections() (*connect.Connections, error) {
	return nil, nil
}

func (c *Config) SaveConnection(connection *connect.Connect) error {
	return nil
}

func (c *Config) UpdateConnection(connection *connect.Connect) (*connect.Connect, error) {
	return nil, nil
}

func (c *Config) DeleteConnection(connection *connect.Connect) error {
	return nil
}
