package kernel

import (
	"errors"

	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/setup"
	"github.com/ssh-connection-manager/kernel/v2/internal/store"
	"github.com/ssh-connection-manager/kernel/v2/pkg/connect"
)

var ErrGetConnectionAtList = errors.New("err get connections")

func List() (*connect.Connections, error) {
	setup.Init()

	connections, err := store.GetConnections()
	if err != nil {
		logger.Error(ErrGetConnectionAtList.Error())
		return nil, ErrGetConnectionAtList
	}

	return connections, nil
}
