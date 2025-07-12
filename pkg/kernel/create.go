package kernel

import (
	"errors"

	"github.com/ssh-connection-manager/kernel/internal/logger"
	"github.com/ssh-connection-manager/kernel/internal/setup"
	"github.com/ssh-connection-manager/kernel/internal/store"
	"github.com/ssh-connection-manager/kernel/pkg/connect"
)

var (
	ErrConnectionByAliasExistsAtCreate = errors.New("connection by alias exists")
	ErrDeletePrivateKeyAtCreate        = errors.New("err delete private key")
	ErrSavePrivateKeyAtCreate          = errors.New("err save private key")
	ErrGetConnectionAtCreate           = errors.New("err get connection")
	ErrSetConnectionAtCreate           = errors.New("err set connection")
)

// Create add connection in file
func Create(connection *connect.Connect) error {
	setup.Init()

	connections, err := store.GetConnections()
	if err != nil {
		logger.Error(ErrGetConnectionAtCreate.Error())
		return ErrGetConnectionAtCreate
	}

	for _, savedConnection := range connections.Connects {
		if savedConnection.Alias == connection.Alias {
			logger.Error(ErrConnectionByAliasExistsAtCreate.Error())
			return ErrConnectionByAliasExistsAtCreate
		}
	}

	if len(connection.SshOptions.PrivateKey) != 0 {
		connection.SshOptions.PrivateKey, err = store.SavePrivateKey(connection)
		if err != nil {
			logger.Error(ErrSavePrivateKeyAtCreate.Error())
			return ErrSavePrivateKeyAtCreate
		}
	}

	connections.Connects = append(connections.Connects, *connection)

	err = store.SetConnections(connections)
	if err != nil {
		if len(connection.SshOptions.PrivateKey) != 0 {
			errDeleteKey := store.DeletePrivateKey(connection)
			if errDeleteKey != nil {
				logger.Error(ErrDeletePrivateKeyAtCreate.Error())
				return ErrDeletePrivateKeyAtCreate
			}
		}

		logger.Error(ErrSetConnectionAtCreate.Error())
		return ErrSetConnectionAtCreate
	}

	return nil
}
