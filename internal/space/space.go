package space

import (
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
)

type Space interface {
	GetConnections() (*connect.Connections, error)
	SaveConnection(connection *connect.Connect) error
	UpdateConnection(connection *connect.Connect) (*connect.Connect, error)
	DeleteConnection(connection *connect.Connect) error
}

var defaultSpace Space

func Get() Space {
	if defaultSpace == nil {
		defaultSpace = &Storage{
			Storage: storage.Get(),
		}
	}

	return defaultSpace
}

func Set(space Space) {
	defaultSpace = space
}
