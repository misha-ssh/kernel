package testutil

import (
	"github.com/ssh-connection-manager/kernel/v2/configs/envconst"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

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
