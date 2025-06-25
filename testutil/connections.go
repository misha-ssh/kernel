package testutil

import (
	"github.com/ssh-connection-manager/kernel/v2/configs/envconst"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

func RemoveFileConnections() error {
	if storage.Exists(storage.GetAppDir(), envconst.FilenameConnections) {
		return storage.Delete(storage.GetAppDir(), envconst.FilenameConnections)
	}

	return nil
}
