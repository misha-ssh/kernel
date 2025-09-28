package main

import (
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

// main for success connect start make command: up-ssh
// kernel.Download this command downloads the remote file
func main() {
	connection := &connect.Connect{
		Alias:     "test",
		Login:     "root",
		Password:  "password",
		Address:   "localhost",
		Type:      connect.TypeSSH,
		CreatedAt: "",
		UpdatedAt: "",
		SshOptions: &connect.SshOptions{
			Port:       22,
			PrivateKey: "",
		},
	}

	remoteFile := "/remote.txt"
	localFile := "~/local.txt"

	err := kernel.Download(connection, remoteFile, localFile)
	if err != nil {
		panic(err)
	}
}
