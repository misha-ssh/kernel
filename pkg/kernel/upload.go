package kernel

import (
	"github.com/misha-ssh/kernel/internal/setup"
	"github.com/misha-ssh/kernel/pkg/connect"
)

func Upload(connection *connect.Connect, localFile string, remoteFile string) error {
	setup.Init()
	return nil
}
