package space

import (
	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/configs/envname"
	"github.com/misha-ssh/kernel/internal/config"
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
)

type Space struct {
	Storage *storage.Storage
}

func New() *Space {
	space := &Space{}

	switch config.Get(envname.Storage) {
	case envconst.TypeLocalStorage:
		space.Storage = storage.NewLocal()
	}

	return space
}

func (s *Space) GetConnections() (*connect.Connections, error) {
	return nil, nil
}

func (s *Space) SaveConnection(connection *connect.Connect) error {
	return nil
}

func (s *Space) UpdateConnection(connection *connect.Connect) (*connect.Connect, error) {
	return nil, nil
}

func (s *Space) DeleteConnection(connection *connect.Connect) error {
	return nil
}
