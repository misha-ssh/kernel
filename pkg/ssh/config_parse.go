package ssh

import (
	"bufio"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/misha-ssh/kernel/pkg/connect"
)

func parseHost(connection *connect.Connect, values []string) error {
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

func parsePort(connection *connect.Connect, values []string) error {
	if strings.ToLower(values[0]) != "port" {
		return nil
	}

	if len(values) < 2 {
		return fmt.Errorf("empty port")
	}

	port, err := strconv.Atoi(values[1])
	if err != nil {
		return err
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

	connection.PrivateKey = values[1]

	return nil
}

func addConnection(connection *connect.Connect, connections *connect.Connections) {
	if !reflect.DeepEqual(connection, new(connect.Connect)) {
		connections.Connects = append(connections.Connects, *connection)
	}

	connection = new(connect.Connect)
}

func parseConnection(s *bufio.Scanner) (*connect.Connections, error) {
	connections := new(connect.Connections)
	connection := new(connect.Connect)

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		fmt.Println(line)

		if strings.HasPrefix(line, "#") {
			continue
		}

		if line == "" {
			addConnection(connection, connections)
			continue
		}

		aliasValues := strings.Split(line, " ")

		for _, err := range []error{
			parseHost(connection, aliasValues),
			parsePort(connection, aliasValues),
			parseLogin(connection, aliasValues),
			parsePrivateKey(connection, aliasValues),
		} {
			if err != nil {
				return nil, err
			}
		}
	}

	fmt.Println(connections)

	if err := s.Err(); err != nil {
		return nil, err
	}

	return connections, nil
}
