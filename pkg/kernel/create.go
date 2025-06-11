package kernel

import (
	"errors"

	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/setup"
	"github.com/ssh-connection-manager/kernel/v2/internal/store"
)

var (
	ErrConnectionByAliasExists = errors.New("connection by alias exists")
	ErrGetConnection           = errors.New("err get connection")
	ErrSetConnection           = errors.New("err set connection")
)

func validateConnection(connections *connect.Connections, connection *connect.Connect) error {
	for _, savedConnection := range connections.Connects {
		if savedConnection.Alias == connection.Alias {
			logger.Error(ErrConnectionByAliasExists.Error())
			return ErrConnectionByAliasExists
		}
	}

	return nil
}

func Create(connection *connect.Connect) error {
	setup.Init()

	connections, err := store.GetConnections()
	if err != nil {
		logger.Error(ErrGetConnection.Error())
		return ErrGetConnection
	}

	err = validateConnection(connections, connection)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	connections.Connects = append(connections.Connects, *connection)

	err = store.SetConnections(connections)
	if err != nil {
		logger.Error(ErrSetConnection.Error())
		return ErrSetConnection
	}

	return nil
}
