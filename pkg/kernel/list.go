package kernel

import (
	"errors"

	"github.com/ssh-connection-manager/kernel/internal/logger"
	"github.com/ssh-connection-manager/kernel/internal/setup"
	"github.com/ssh-connection-manager/kernel/internal/store"
	"github.com/ssh-connection-manager/kernel/pkg/connect"
)

var ErrGetConnectionAtList = errors.New("err get connections")

// List get list with connections from file
func List() (*connect.Connections, error) {
	setup.Init()

	connections, err := store.GetConnections()
	if err != nil {
		logger.Error(ErrGetConnectionAtList.Error())
		return nil, ErrGetConnectionAtList
	}

	return connections, nil
}
