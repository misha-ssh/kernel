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

type Ssh struct{}

func NewSshConnector() *Ssh {
	return &Ssh{}
}

// NewSession establishes a new SSH session with the remote server.
// It handles the complete connection lifecycle including:
// - Authentication (password or private key)
// - Client creation
// - Terminal setup
// - Error handling and resource cleanup
// Returns an active SSH session or error if any step fails.
func (s *Ssh) NewSession(connection *Connect) (*ssh.Session, error) {
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
		if errClient := client.Close(); errClient != nil {
			logger.Error(errClient.Error())
			return nil, errClient
		}

		logger.Error(err.Error())
		return nil, err
	}

	err = createTerminalSession(session)
	if err != nil {
		if errClient := client.Close(); errClient != nil {
			logger.Error(errClient.Error())
			return nil, errClient
		}
		if errSession := session.Close(); errSession != nil {
			logger.Error(errSession.Error())
			return nil, errSession
		}

		logger.Error(err.Error())
		return nil, err
	}

	return session, nil
}

// Connect starts an interactive shell session using the established SSH connection.
// It manages the session lifecycle including proper cleanup on exit.
// Returns error if shell startup or session wait fails.
func (s *Ssh) Connect(session *ssh.Session) error {
	defer func() {
		if err := session.Close(); err != nil {
			logger.Error(err.Error())
		}
	}()

	err := session.Shell()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = session.Wait()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
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
