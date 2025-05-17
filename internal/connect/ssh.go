package connect

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
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

func auth(connection *Connect) []ssh.AuthMethod {
	if len(connection.SshOptions.PrivateKey) > 1 {
		return []ssh.AuthMethod{
			ssh.Password(connection.Password),
		}
	}
	//todo доделать авторизацию через ключ
	key := ssh.ParsePrivateKey(connection.SshOptions.PrivateKey)

	return []ssh.AuthMethod{
		ssh.PublicKeys(connection.SshOptions.PrivateKey),
	}
}

func getClientConfig(connection *Connect) *ssh.ClientConfig {
	callback := ssh.InsecureIgnoreHostKey()
	sshAuth := auth(connection)

	return &ssh.ClientConfig{
		Timeout:         Timeout,
		User:            connection.Login,
		Auth:            sshAuth,
		HostKeyCallback: callback,
	}
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

func (s *SshConnect) Connect(connection *Connect) error {
	config := getClientConfig(connection)
	hostWithPort := net.JoinHostPort(
		connection.Address,
		fmt.Sprint(connection.SshOptions.Port),
	)

	client, err := getClient(hostWithPort, config)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	session, err := getSession(client)
	if err != nil {
		client.Close()
		logger.Error(err.Error())
		return err
	}

	err = createTerminalSession(session)
	if err != nil {
		client.Close()
		session.Close()
		logger.Error(err.Error())
		return err
	}

	err = session.Shell()
	if err != nil {
		session.Close()
		client.Close()
		logger.Error(err.Error())
		return err
	}

	err = session.Wait()
	if err != nil {
		client.Close()
		session.Close()
		logger.Error(err.Error())
		return err
	}

	return nil
}
