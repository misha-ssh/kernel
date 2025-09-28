package kernel

import (
	"io"
	"os"

	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/setup"
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/pkg/sftp"
)

func Upload(connection *connect.Connect, uploadLocalFile string, uploadRemoteFile string) error {
	setup.Init()

	localPath, localFilename := storage.GetDirectionAndFilename(uploadLocalFile)
	localFile, err := storage.GetOpenFile(localPath, localFilename, os.O_RDWR)
	if err != nil {
		return err
	}

	client, err := connect.NewSftp(connection)
	if err != nil {
		return err
	}
	defer func(client *sftp.Client) {
		err = client.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(client)

	remoteFile, err := client.Create(uploadRemoteFile)
	if err != nil {
		return err
	}
	defer func(remote *sftp.File) {
		err = remote.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(remoteFile)

	_, err = io.Copy(remoteFile, localFile)
	return err
}
