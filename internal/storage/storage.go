package storage

import (
	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/configs/envname"
	"os"
	"os/user"
	"path/filepath"
)

type Storage interface {
	Create(filename string) error
	Delete(filename string) error
	Exists(filename string) bool
	Get(filename string) (string, error)
	Write(filename string, data string) error
	GetOpenFile(filename string, flags int) (*os.File, error)
}

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

// GetDirSSH get dir ssh
func GetDirSSH() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(homeDir, envconst.DirectionsUserPrivateKey)
}

// GetPrivateKeysDir get dir where save private keys
func GetPrivateKeysDir() string {
	return filepath.Join(GetAppDir(), envconst.DirectionPrivateKeys)
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
