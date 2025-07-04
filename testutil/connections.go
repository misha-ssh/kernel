package testutil

import (
	"github.com/ssh-connection-manager/kernel/v2/configs/envconst"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

// RemoveFileConnections delete file connection
func RemoveFileConnections() error {
	if storage.Exists(storage.GetAppDir(), envconst.FilenameConnections) {
		return storage.Delete(storage.GetAppDir(), envconst.FilenameConnections)
	}

	return nil
}
