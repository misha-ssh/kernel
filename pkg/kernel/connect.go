package kernel

import (
	"errors"

	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/setup"
	"github.com/misha-ssh/kernel/pkg/connect"
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
		ssh := &connect.Ssh{}
		session, err := ssh.Session(connection)
		if err != nil {
			logger.Error(ErrSshSession)
			return ErrSshSession
		}

		err = ssh.Connect(session)
		if err != nil {
			logger.Error(ErrSshConnect)
			return ErrSshConnect
		}
	default:
		return ErrTypeConnect
	}

	return nil
}
