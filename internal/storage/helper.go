package storage

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/configs/envname"
)

const CharHidden = "."

// GetAppDir get dir application
func GetAppDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	hiddenDir := CharHidden + envconst.AppName

	if os.Getenv(envname.Testing) == envconst.IsTesting {
		return filepath.Join(os.TempDir(), hiddenDir)
	}

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
