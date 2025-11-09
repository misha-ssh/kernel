package kernel

import (
	"errors"

	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/setup"
	"github.com/misha-ssh/kernel/internal/store"
	"github.com/misha-ssh/kernel/pkg/connect"
)

var ErrGetConnectionAtList = errors.New("err get connections")

// List get list with connections from file
func List() (*connect.Connections, error) {
	setup.Init()

	connections, err := store.GetConnections()
	if err != nil {
		logger.Error(err.Error())
		return nil, ErrGetConnectionAtList
	}

	return connections, nil
}
