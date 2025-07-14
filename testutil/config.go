package testutil

import (
	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/internal/storage"
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
