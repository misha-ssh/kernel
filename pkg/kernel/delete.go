package kernel

import (
	"errors"

	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/setup"
	"github.com/ssh-connection-manager/kernel/v2/internal/store"
)

var (
	ErrNotFoundConnectionAtDelete = errors.New("err found connection")
	ErrGetConnectionAtDelete      = errors.New("err get connection")
	ErrSetConnectionAtDelete      = errors.New("err set connection")
)

func Delete(connection *connect.Connect) error {
	setup.Init()

	connections, err := store.GetConnections()
	if err != nil {
		logger.Error(ErrGetConnectionAtDelete.Error())
		return ErrGetConnectionAtDelete
	}

	for key, savedConnection := range connections.Connects {
		if savedConnection.Alias == connection.Alias {
			connections.Connects = append(connections.Connects[:key], connections.Connects[key+1:]...)

			err = store.SetConnections(connections)
			if err != nil {
				logger.Error(ErrSetConnectionAtDelete.Error())
				return ErrSetConnectionAtDelete
			}

			if len(connection.SshOptions.PrivateKey) != 0 {
				return store.DeletePrivateKey(connection)
			}
		}
	}

	return ErrNotFoundConnectionAtDelete
}
