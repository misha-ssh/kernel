package storage

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/ssh-connection-manager/kernel/v2/config/envconst"
)

type Storage interface {
	Exists(filename string) bool
	Create(filename string) error
	Get(filename string) (string, error)
	Delete(filename string) error
	Write(filename string, data string) error
	GetOpenFile(filename string, flags int) (*os.File, error)
}

func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	hiddenDir := "." + envconst.AppName

	return filepath.Join(usr.HomeDir, hiddenDir)
}
