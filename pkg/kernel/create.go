package kernel

import (
	"errors"

	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/setup"
	"github.com/ssh-connection-manager/kernel/v2/internal/store"
)

var (
	ErrConnectionByAliasExistsAtCreate = errors.New("connection by alias exists")
	ErrGetConnectionAtCreate           = errors.New("err get connection")
	ErrSetConnectionAtCreate           = errors.New("err set connection")
)

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

	connections.Connects = append(connections.Connects, *connection)

	err = store.SetConnections(connections)
	if err != nil {
		logger.Error(ErrSetConnectionAtCreate.Error())
		return ErrSetConnectionAtCreate
	}

	return nil
}
