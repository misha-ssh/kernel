package ssh

import (
	"bufio"
	"fmt"
	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"os"
	"strings"
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
	var connections connect.Connections

	file, err := c.LocalStorage.GetOpenFile(envconst.FilenameConfigSSH, os.O_RDWR)
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(file)

	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		if strings.HasPrefix(line, "#") {
			continue
		}

		if strings.Contains(line, "Host") {
			value := strings.Split(line, " ")

			if value[0] == "Host" {
				connection := new(connect.Connect)
				connection.Alias = line[5:]

				connections.Connects = append(connections.Connects, *connection)
			}
		}

		fmt.Println(line)
	}

	fmt.Println(connections)

	if err = s.Err(); err != nil {
		fmt.Println("err: ", err)
	}

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
