package connect

import (
	"strings"

	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/storage"
	"golang.org/x/crypto/ssh"
)

func parsePrivateKey(keyName string, passphrase string) (ssh.Signer, error) {
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

		key, err = ssh.ParsePrivateKeyWithPassphrase(dataSshKey, []byte(passphrase))
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
	}

	return key, nil
}
