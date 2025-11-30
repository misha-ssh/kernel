package ssh

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/pkg/connect"
)

func parseAlias(connection *connect.Connect, values []string) error {
	if strings.ToLower(values[0]) != "host" {
		return nil
	}

	if len(values) < 2 {
		return fmt.Errorf("empty host")
	}

	if strings.Contains(values[1], "*") || strings.Contains(values[1], "!") {
		return nil
	}

	connection.Alias = values[1]
	return nil
}

func parseAddress(connection *connect.Connect, values []string) error {
	if strings.ToLower(values[0]) != "hostname" {
		return nil
	}

	if len(values) < 2 {
		return fmt.Errorf("empty hostname")
	}

	connection.Address = values[1]
	return nil
}

func parsePort(connection *connect.Connect, values []string) error {
	if strings.ToLower(values[0]) != "port" {
		return nil
	}

	if len(values) < 2 {
		return fmt.Errorf("empty port")
	}

	port, err := strconv.Atoi(values[1])
	if err != nil {
		return fmt.Errorf("invalid port: %q", values[1])
	}

	connection.Port = port
	return nil
}

func parseLogin(connection *connect.Connect, values []string) error {
	if strings.ToLower(values[0]) != "user" {
		return nil
	}

	if len(values) < 2 {
		return fmt.Errorf("empty user")
	}

	connection.Login = values[1]
	return nil
}

func parsePrivateKey(connection *connect.Connect, values []string) error {
	if strings.ToLower(values[0]) != "identityfile" {
		return nil
	}

	if len(values) < 2 {
		return fmt.Errorf("empty private key")
	}

	pathKey := values[1]

	if strings.HasPrefix(pathKey, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		pathKey = filepath.Join(home, pathKey[2:])
	}

	connection.PrivateKey = pathKey
	return nil
}

func isComment(line string) bool {
	return strings.HasPrefix(line, "#")
}

func isEmptyLine(line string) bool {
	return line == ""
}

func parseLine(connection *connect.Connect, line string) error {
	values := strings.Fields(line)
	if len(values) == 0 {
		return nil
	}

	parsers := []func(*connect.Connect, []string) error{
		parseAlias,
		parseAddress,
		parseLogin,
		parsePort,
		parsePrivateKey,
	}

	for _, parser := range parsers {
		if err := parser(connection, values); err != nil {
			return err
		}
	}

	return nil
}

func saveCurrentConnection(connection *connect.Connect, connections *connect.Connections) {
	if connection != nil && !isConnectionEmpty(connection) {
		if connection.Port == 0 {
			connection.Port = 22
		}

		if err := connection.Validate(); err == nil {
			connections.Connects = append(connections.Connects, *connection)
		}
	}

	*connection = connect.Connect{}
}

func isConnectionEmpty(connection *connect.Connect) bool {
	return connection.Alias == "" &&
		connection.Port == 0 &&
		connection.Login == "" &&
		connection.Address == "" &&
		connection.PrivateKey == ""
}

func parseConnections(file *os.File) (*connect.Connections, error) {
	defer func() {
		err := file.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}()

	s := bufio.NewScanner(file)

	connections := new(connect.Connections)
	current := new(connect.Connect)

	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		switch {
		case isComment(line):
			continue
		case isEmptyLine(line):
			saveCurrentConnection(current, connections)
		default:
			if err := parseLine(current, line); err != nil {
				return nil, err
			}
		}
	}

	saveCurrentConnection(current, connections)

	if err := s.Err(); err != nil {
		return nil, err
	}

	return connections, nil
}

func prepareConnection(connection *connect.Connect) string {
	var configConnection string

	configConnection += fmt.Sprintf("\n\nHost %v", connection.Alias)
	configConnection += fmt.Sprintf("\n\tHostName %v", connection.Address)
	configConnection += fmt.Sprintf("\n\tUser %v", connection.Login)
	configConnection += fmt.Sprintf("\n\tPort %v", strconv.Itoa(connection.Port))
	configConnection += fmt.Sprintf("\n\tIdentityFile %v", connection.PrivateKey)

	return configConnection
}

func addConnection(connection *connect.Connect, file *os.File) error {
	defer func() {
		err := file.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}()

	if err := connection.Validate(); err != nil {
		return err
	}

	_, err := file.WriteString(prepareConnection(connection))
	return err
}
