package ssh

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
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

	connection := new(connect.Connect)

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		fmt.Println(line)

		if strings.HasPrefix(line, "#") {
			continue
		}

		value := strings.Split(line, " ")

		if line == "" {
			if !reflect.DeepEqual(connection, &connect.Connect{}) {
				connections.Connects = append(connections.Connects, *connection)
			}

			connection = new(connect.Connect)
		}

		if value[0] == "Host" {
			if len(value) < 1 {
				continue
			}

			if strings.Contains(value[1], "*") || strings.Contains(value[1], "!") {
				continue
			}

			connection.Alias = value[1]
		}

		if value[0] == "HostName" {
			connection.Address = value[1]
		}

		if value[0] == "User" {
			connection.Login = value[1]
		}

		if value[0] == "Port" {
			port, err := strconv.Atoi(value[1])
			if err != nil {
				return nil, err
			}
			connection.Port = port
		}

		if value[0] == "IdentityFile" {
			connection.PrivateKey = value[1]
		}
	}

	fmt.Println(connections)

	if err = s.Err(); err != nil {
		return nil, err
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
