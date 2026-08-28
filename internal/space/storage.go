package space

import (
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
)

type Storage struct {
	Storage storage.Storage
}

func (s *Storage) GetConnections() (*connect.Connections, error) {
	return nil, nil
}

func (s *Storage) SaveConnection(connection *connect.Connect) error {
	return nil
}

func (s *Storage) UpdateConnection(connection *connect.Connect) (*connect.Connect, error) {
	return nil, nil
}

func (s *Storage) DeleteConnection(connection *connect.Connect) error {
	return nil
}
