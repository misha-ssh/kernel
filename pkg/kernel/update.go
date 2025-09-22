package kernel

import (
	"errors"

	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/setup"
	"github.com/misha-ssh/kernel/internal/store"
	"github.com/misha-ssh/kernel/pkg/connect"
)

var (
	ErrNotFoundConnectionAtUpdate = errors.New("err not found connection")
	ErrSavePrivateKeyAtUpdate     = errors.New("err save private key")
	ErrGetConnectionAtUpdate      = errors.New("err get connection")
	ErrSetConnectionAtUpdate      = errors.New("err set connection")
)

// Update change connection by alias
func Update(connection *connect.Connect, oldAlias string) error {
	setup.Init()

	err := connection.Validate()
	if err != nil {
		return err
	}

	connections, err := store.GetConnections()
	if err != nil {
		logger.Error(ErrGetConnectionAtUpdate.Error())
		return ErrGetConnectionAtUpdate
	}

	for i, savedConnection := range connections.Connects {
		if savedConnection.Alias == oldAlias {
			connection.SshOptions.PrivateKey, err = store.UpdatePrivateKey(connection)
			if err != nil {
				logger.Error(ErrSavePrivateKeyAtUpdate.Error())
				return ErrSavePrivateKeyAtUpdate
			}

			connections.Connects[i] = *connection

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
