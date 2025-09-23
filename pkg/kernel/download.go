package kernel

import (
	"github.com/misha-ssh/kernel/internal/setup"
	"github.com/misha-ssh/kernel/pkg/connect"
)

func Download(connection *connect.Connect, remoteFile string, localFile string) error {
	setup.Init()
	return nil
}
