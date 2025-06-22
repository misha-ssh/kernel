package connect

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"golang.org/x/crypto/ssh"
)

const (
	TypeConnect = "tcp"

	Timeout = 20 * time.Second

	DisableEcho = 0
	IgnoreCR    = 1

	TypeTerm       = "vt100"
	HeightTerminal = 80
	WidthTerminal  = 40
)

type SshConnect struct{}

func NewSshConnect() *SshConnect {
	return &SshConnect{}
}

func auth(connection *Connect) ([]ssh.AuthMethod, error) {
	if len(connection.SshOptions.PrivateKey) == 0 {
		return []ssh.AuthMethod{
			ssh.Password(connection.Password),
		}, nil
	}

	direction, filename := storage.GetDirectionAndFilename(connection.SshOptions.PrivateKey)
	dataPrivateKey, err := storage.Get(direction, filename)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	key, err := ssh.ParsePrivateKey([]byte(dataPrivateKey))
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return []ssh.AuthMethod{
		ssh.PublicKeys(key),
	}, nil
}

func getClientConfig(connection *Connect) (*ssh.ClientConfig, error) {
	callback := ssh.InsecureIgnoreHostKey()
	sshAuth, err := auth(connection)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return &ssh.ClientConfig{
		Timeout:         Timeout,
		User:            connection.Login,
		Auth:            sshAuth,
		HostKeyCallback: callback,
	}, nil
}

func createTerminalSession(session *ssh.Session) error {
	modes := ssh.TerminalModes{
		ssh.ECHO:  DisableEcho,
		ssh.IGNCR: IgnoreCR,
	}

	err := session.RequestPty(TypeTerm, HeightTerminal, WidthTerminal, modes)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	return nil
}

func getClient(hostWithPort string, config *ssh.ClientConfig) (*ssh.Client, error) {
	client, err := ssh.Dial(TypeConnect, hostWithPort, config)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return client, nil
}

func getSession(client *ssh.Client) (*ssh.Session, error) {
	session, err := client.NewSession()
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return session, nil
}

func (s *SshConnect) Connect(connection *Connect) (*ssh.Session, error) {
	config, err := getClientConfig(connection)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	hostWithPort := net.JoinHostPort(
		connection.Address,
		fmt.Sprint(connection.SshOptions.Port),
	)

	client, err := getClient(hostWithPort, config)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	session, err := getSession(client)
	if err != nil {
		client.Close()
		logger.Error(err.Error())
		return nil, err
	}

	err = createTerminalSession(session)
	if err != nil {
		client.Close()
		session.Close()
		logger.Error(err.Error())
		return nil, err
	}

	return session, nil
}
