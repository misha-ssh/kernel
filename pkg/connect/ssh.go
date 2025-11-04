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

type Ssh struct {
	Connection *Connect
}

// Session establishes a new SSH session with the remote server
func (s *Ssh) Session() (*ssh.Session, error) {
	client, err := s.Client()
	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
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

// Connect starts an interactive shell session using the established SSH connection
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

// Client create ssh client from config and Auth
func (s *Ssh) Client() (*ssh.Client, error) {
	sshAuth, err := s.Auth()
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		Timeout:         Timeout,
		User:            s.Connection.Login,
		Auth:            sshAuth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	hostWithPort := net.JoinHostPort(
		s.Connection.Address,
		fmt.Sprint(s.Connection.SshOptions.Port),
	)

	return ssh.Dial("tcp", hostWithPort, config)
}

// Auth automate defines method auth from Connect
func (s *Ssh) Auth() ([]ssh.AuthMethod, error) {
	var authMethod []ssh.AuthMethod

	if len(s.Connection.Password) > 0 {
		authMethod = append(authMethod, ssh.Password(s.Connection.Password))
	}

	if len(s.Connection.SshOptions.PrivateKey) > 0 {
		key, err := parsePrivateKey(s.Connection.SshOptions.PrivateKey)
		if err != nil {
			return nil, err
		}

		authMethod = append(authMethod, ssh.PublicKeys(key))
	}

	if len(s.Connection.Password) == 0 && len(s.Connection.SshOptions.PrivateKey) == 0 {
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
