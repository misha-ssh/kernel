package ssh

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/misha-ssh/kernel/pkg/connect"
)

type ConfigParser struct {
	connections *connect.Connections
	current     *connect.Connect
}

func parseAlias(connection *connect.Connect, values []string) error {
	if strings.ToLower(values[0]) != "host" {
		return nil
	}

	if len(values) < 2 {
		return fmt.Errorf("empty host")
	}

	//todo add logic for pase host with * and ! for set data
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

func (p *ConfigParser) parseLine(line string) error {
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
		if err := parser(p.current, values); err != nil {
			return err
		}
	}

	return nil
}

func (p *ConfigParser) saveCurrentConnection() {
	if p.current != nil && !p.isConnectionEmpty() {
		if p.current.Port == 0 {
			p.current.Port = 22
		}

		p.connections.Connects = append(p.connections.Connects, *p.current)
	}

	p.current = &connect.Connect{}
}

func (p *ConfigParser) isConnectionEmpty() bool {
	return p.current.Alias == "" &&
		p.current.Port == 0 &&
		p.current.Login == "" &&
		p.current.PrivateKey == ""
}

func parseConnections(s *bufio.Scanner) (*connect.Connections, error) {
	parser := &ConfigParser{
		connections: &connect.Connections{},
		current:     &connect.Connect{},
	}

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		fmt.Println(line)

		switch {
		case isComment(line):
			continue
		case isEmptyLine(line):
			parser.saveCurrentConnection()
		default:
			if err := parser.parseLine(line); err != nil {
				return nil, err
			}
		}
	}

	parser.saveCurrentConnection()

	if err := s.Err(); err != nil {
		return nil, err
	}

	fmt.Println(parser.connections)
	return parser.connections, nil
}
