package storage

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"strings"

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

// GetUserPrivateKey get file with ssh keys
func GetUserPrivateKey() ([]string, error) {
	var privateKeys []string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	listDeniedPatternKeys := []string{
		".pub",
		"known_hosts",
		"config",
		"authorized_keys",
	}

	keysDir := filepath.Join(homeDir, envconst.DirectionsUserPrivateKey)

	keys, err := os.ReadDir(keysDir)
	if err != nil || len(keys) == 0 {
		return []string{}, errors.New("cannot find user private keys")
	}

	for _, key := range keys {
		if key.IsDir() {
			continue
		}

		keyName := key.Name()

		containsPattern := false
		for _, pattern := range listDeniedPatternKeys {
			if strings.Contains(keyName, pattern) {
				containsPattern = true
				break
			}
		}

		if !containsPattern {
			privateKeys = append(privateKeys, filepath.Join(keysDir, keyName))
		}
	}

	if len(privateKeys) == 0 {
		return []string{}, errors.New("cannot find user private keys")
	}

	return privateKeys, nil
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
