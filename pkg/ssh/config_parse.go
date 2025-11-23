package ssh

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/misha-ssh/kernel/pkg/connect"
)

func parseAlias(connection *connect.Connect, values []string) error {
	if strings.ToLower(values[0]) != "host" {
		return nil
	}

	if len(values) < 2 {
		return fmt.Errorf("empty host")
	}

	//todo add logic for parse host with * and ! for set data
	//todo this is operation used in last order
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

	connection.PrivateKey = values[1]
	return nil
}

func isComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "#")
}

func isEmptyLine(line string) bool {
	return strings.TrimSpace(line) == ""
}

func parseLine(connection *connect.Connect, line string) error {
	values := strings.Fields(line)
	if len(values) == 0 {
		return nil
	}

	parsers := []func(*connect.Connect, []string) error{
		parseAlias,
		parsePort,
		parseLogin,
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

		connections.Connects = append(connections.Connects, *connection)
	}

	*connection = connect.Connect{}
}

func isConnectionEmpty(connection *connect.Connect) bool {
	return connection.Alias == "" &&
		connection.Port == 0 &&
		connection.Login == "" &&
		connection.PrivateKey == ""
}

func parseConnections(s *bufio.Scanner) (*connect.Connections, error) {
	connections := new(connect.Connections)
	current := new(connect.Connect)

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		fmt.Println(line)

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

	fmt.Println(connections)
	return connections, nil
}
