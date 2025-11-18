package ssh

import (
	"strings"

	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/storage"
	"golang.org/x/crypto/ssh"
)

func (s *SSH) parsePrivateKey() (ssh.Signer, error) {
	currentStorage := storage.Get()

	data, err := currentStorage.Get(s.Connection.PrivateKey)
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

		key, err = ssh.ParsePrivateKeyWithPassphrase(dataSshKey, []byte(s.Connection.Passphrase))
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
	}

	return key, nil
}
