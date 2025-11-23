package ssh

import (
	"bufio"
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

	_, _ = parseConnections(bufio.NewScanner(file))

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
