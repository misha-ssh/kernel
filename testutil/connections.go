package testutil

import (
	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/internal/storage"
)

// RemoveFileConnections delete file connection
func RemoveFileConnections() error {
	if storage.Exists(storage.GetAppDir(), envconst.FilenameConnections) {
		return storage.Delete(storage.GetAppDir(), envconst.FilenameConnections)
	}

	return nil
}
