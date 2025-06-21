package storage

import (
	"os/user"
	"path/filepath"

	"github.com/ssh-connection-manager/kernel/v2/configs/envconst"
)

func GetAppDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	hiddenDir := "." + envconst.AppName

	return filepath.Join(usr.HomeDir, hiddenDir)
}

func GetPrivateKeysDir() string {
	appDir := GetAppDir()

	return filepath.Join(appDir, envconst.DirectionPrivateKeys)
}
