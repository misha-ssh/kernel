package ssh

import (
	"github.com/misha-ssh/kernel/internal/storage"
)

type Config struct {
	Path string
}

func NewConfig() *Config {
	return &Config{
		Path: storage.GetDirSSH(),
	}
}
