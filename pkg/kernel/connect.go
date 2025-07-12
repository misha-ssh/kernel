package kernel

import (
	"errors"

	"github.com/ssh-connection-manager/kernel/internal/logger"
	"github.com/ssh-connection-manager/kernel/internal/setup"
	"github.com/ssh-connection-manager/kernel/pkg/connect"
)

var (
	ErrSshSession  = errors.New("err ssh session")
	ErrSshConnect  = errors.New("err ssh connect")
	ErrTypeConnect = errors.New("err type connect")
)

// Connect from connection by type connection
func Connect(connection *connect.Connect) error {
	setup.Init()

	switch connection.Type {
	case connect.TypeSSH:
		sshConnector := &connect.Ssh{}
		session, err := sshConnector.Session(connection)
		if err != nil {
			logger.Error(ErrSshSession)
			return ErrSshSession
		}

		err = sshConnector.Connect(session)
		if err != nil {
			logger.Error(ErrSshConnect)
			return ErrSshConnect
		}
	default:
		return ErrTypeConnect
	}

	return nil
}
