package config

import "github.com/ssh-connection-manager/kernel/v2/internal/init"

func Init() error {
	err := init.FileConfig()
	if err != nil {
		return err
	}

	return nil
}
