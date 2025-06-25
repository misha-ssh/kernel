package kernel

import (
	"errors"

	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/setup"
	"golang.org/x/crypto/ssh"
)

var (
	ErrSshConnect  = errors.New("err ssh connect")
	ErrTypeConnect = errors.New("err type connect")
)

func Connect(connection *connect.Connect) (*ssh.Session, error) {
	setup.Init()

	switch connection.Type {
	case connect.TypeSSH:
		sshConnect := connect.NewSshConnect()
		session, err := sshConnect.Connect(connection)
		if err != nil {
			logger.Error(ErrSshConnect.Error())
			return nil, ErrSshConnect
		}
		return session, nil
	default:
		return nil, ErrTypeConnect
	}
}
