package connect

import (
	"fmt"
	"net"
	"os"

	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/storage"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

const (
	TypeConnect = "tcp"

	Timeout = 0

	EnableMod = 1
	ICRNLMod  = 1
	INLCRMod  = 1
	ISIGMod   = 1

	ISPEED = 115200
	OSPEED = 115200

	TypeTerm       = "xterm-256color"
	HeightTerminal = 80
	WidthTerminal  = 40
)

type Ssh struct{}

// Session establishes a new SSH session with the remote server.
// It handles the complete connection lifecycle including:
// - Authentication (password or private key)
// - Client creation
// - Terminal setup
// - Error handling and resource cleanup
// Returns an active SSH session or error if any step fails.
func (s *Ssh) Session(connection *Connect) (*ssh.Session, error) {
	client, err := s.Client(connection)
	if err != nil {
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

	fd := int(os.Stdin.Fd())

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer func() {
		if err = term.Restore(fd, oldState); err != nil {
			logger.Error(err.Error())
		}
	}()

	err = session.Shell()
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

func (s *Ssh) Client(connection *Connect) (*ssh.Client, error) {
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

	return client, nil
}

func auth(connection *Connect) ([]ssh.AuthMethod, error) {
	var authMethod []ssh.AuthMethod

	if len(connection.Password) > 0 {
		authMethod = append(authMethod, ssh.Password(connection.Password))
	}

	if len(connection.SshOptions.PrivateKey) > 0 {
		key, err := parsePrivateKey(connection.SshOptions.PrivateKey)
		if err != nil {
			return nil, err
		}

		authMethod = append(authMethod, ssh.PublicKeys(key))
	}

	if len(connection.Password) == 0 && len(connection.SshOptions.PrivateKey) == 0 {
		userPrivateKeys, err := storage.GetUserPrivateKey()
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}

		var successKeys []ssh.Signer

		for _, privateKey := range userPrivateKeys {
			key, err := parsePrivateKey(privateKey)
			if err != nil {
				logger.Error(err.Error())
				continue
			}

			successKeys = append(successKeys, key)
		}

		if len(successKeys) == 0 {
			return nil, fmt.Errorf("no authentication methods available")
		}

		authMethod = append(authMethod, ssh.PublicKeys(successKeys...))
	}

	return authMethod, nil
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
	fd := int(os.Stdin.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		width = WidthTerminal
		height = HeightTerminal
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          EnableMod,
		ssh.ICRNL:         ICRNLMod,
		ssh.INLCR:         INLCRMod,
		ssh.ISIG:          ISIGMod,
		ssh.TTY_OP_ISPEED: ISPEED,
		ssh.TTY_OP_OSPEED: OSPEED,
	}

	if err = session.RequestPty(TypeTerm, height, width, modes); err != nil {
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
