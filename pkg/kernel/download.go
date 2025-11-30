package kernel

import (
	sftp2 "github.com/misha-ssh/kernel/pkg/sftp"
	"io"
	"os"

	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/setup"
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/pkg/sftp"
)

func Download(connection *connect.Connect, downloadRemoteFile string, downloadLocalFile string) error {
	setup.Init()

	sp := sftp2.Sftp{
		Connection: connection,
	}

	client, err := sp.Client()
	if err != nil {
		return err
	}
	defer func(client *sftp.Client) {
		err = client.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(client)

	remoteFile, err := client.Open(downloadRemoteFile)
	if err != nil {
		return err
	}
	defer func(remote *sftp.File) {
		err = remote.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}(remoteFile)

	localPath, localFilename := storage.GetDirectionAndFilename(downloadLocalFile)
	err = storage.Create(localPath, localFilename)
	if err != nil {
		return err
	}

	localFile, err := storage.GetOpenFile(localPath, localFilename, os.O_RDWR)
	if err != nil {
		return err
	}

	if _, err = io.Copy(localFile, remoteFile); err != nil {
		return err
	}

	return localFile.Sync()
}
