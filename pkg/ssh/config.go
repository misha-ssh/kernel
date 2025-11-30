package ssh

import (
	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"os"
)

type Config struct {
	LocalStorage *storage.Local
}

func NewConfig() *Config {
	return &Config{
		LocalStorage: &storage.Local{
			Path: storage.GetDirSSH(),
		},
	}
}

func (c *Config) GetConnections() (*connect.Connections, error) {
	file, err := c.LocalStorage.GetOpenFile(envconst.FilenameConfigSSH, os.O_RDWR)
	if err != nil {
		return nil, err
	}

	return parseConnections(file)
}

func (c *Config) SaveConnection(connection *connect.Connect) error {
	file, err := c.LocalStorage.GetOpenFile(envconst.FilenameConfigSSH, os.O_WRONLY|os.O_APPEND|os.O_CREATE)
	if err != nil {
		return err
	}

	return addConnection(connection, file)
}

func (c *Config) UpdateConnection(connection *connect.Connect) (*connect.Connect, error) {
	return nil, nil
}

func (c *Config) DeleteConnection(connection *connect.Connect) error {
	return nil
}
