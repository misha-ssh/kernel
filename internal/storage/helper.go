package storage

import (
	"os/user"
	"path/filepath"

	"github.com/ssh-connection-manager/kernel/v2/configs/envconst"
)

const CharHidden = "."

// GetAppDir get dir application
func GetAppDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	hiddenDir := CharHidden + envconst.AppName

	return filepath.Join(usr.HomeDir, hiddenDir)
}

// GetPrivateKeysDir get dir where save private keys
func GetPrivateKeysDir() string {
	appDir := GetAppDir()

	return filepath.Join(appDir, envconst.DirectionPrivateKeys)
}

// GetDirectionAndFilename get dir and filename from full path
func GetDirectionAndFilename(fullPath string) (string, string) {
	return filepath.Dir(fullPath),
		filepath.Base(fullPath)
}

// GetFullPath get full path from dir and filename
func GetFullPath(direction string, filename string) string {
	return filepath.Join(direction, filename)
}
