package connect

import (
	"fmt"
	"strings"
	"syscall"

	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/storage"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func parsePrivateKeyWithPass(keyName string, dataKey []byte) (ssh.Signer, error) {
	fmt.Printf("Enter passphrase for %v (ctrl+m for skip): ", keyName)

	passphrase, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n")

	return ssh.ParsePrivateKeyWithPassphrase(dataKey, passphrase)
}

func parsePrivateKey(keyName string) (ssh.Signer, error) {
	direction, filename := storage.GetDirectionAndFilename(keyName)
	data, err := storage.Get(direction, filename)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	dataSshKey := []byte(data)

	key, err := ssh.ParsePrivateKey(dataSshKey)
	if err != nil {
		if !strings.Contains(err.Error(), "passphrase") {
			logger.Error(err.Error())
			return nil, err
		}

		key, err = parsePrivateKeyWithPass(filename, dataSshKey)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
	}

	return key, nil
}
