package kernel

import (
	"errors"

	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/setup"
	"github.com/ssh-connection-manager/kernel/v2/internal/store"
)

var (
	ErrNotFoundConnectionAtUpdate = errors.New("err not found connection")
	ErrGetConnectionAtUpdate      = errors.New("err get connection")
	ErrSetConnectionAtUpdate      = errors.New("err set connection")
)

func Update(connection *connect.Connect, oldAlias string) error {
	setup.Init()

	connections, err := store.GetConnections()
	if err != nil {
		logger.Error(ErrGetConnectionAtUpdate.Error())
		return ErrGetConnectionAtUpdate
	}

	for _, savedConnection := range connections.Connects {
		if savedConnection.Alias == oldAlias {
			connection.SshOptions.PrivateKey, err = store.UpdatePrivateKey(connection)
			if err != nil {
				logger.Error(ErrSavePrivateKeyAtCreate.Error())
				return ErrSavePrivateKeyAtCreate
			}

			connections.Connects = append(connections.Connects, *connection)

			err = store.SetConnections(connections)
			if err != nil {
				logger.Error(ErrSetConnectionAtUpdate.Error())
				return ErrSetConnectionAtUpdate
			}

			return nil
		}
	}

	return ErrNotFoundConnectionAtUpdate
}
