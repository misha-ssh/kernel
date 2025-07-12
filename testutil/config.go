package testutil

import (
	"github.com/ssh-connection-manager/kernel/configs/envconst"
	"github.com/ssh-connection-manager/kernel/internal/storage"
)

// CreateFileConfig create file config for kernel
func CreateFileConfig() error {
	filename := envconst.FilenameConfig
	direction := storage.GetAppDir()

	if !storage.Exists(direction, filename) {
		err := storage.Create(direction, filename)
		if err != nil {
			return err
		}
	}

	return nil
}
