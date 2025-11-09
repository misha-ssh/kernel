package main

import (
	"time"

	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

// main for success connect start make command: up-ssh
// kernel.Upload this command upload the local file
func main() {
	connection := &connect.Connect{
		Alias:     "test",
		Login:     "root",
		Password:  "password",
		Address:   "localhost",
		Type:      connect.TypeSSH,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
		SshOptions: &connect.SshOptions{
			Port:       22,
			PrivateKey: "",
			Passphrase: "",
		},
	}

	remoteFile := "/upload.txt"
	localFile := "/absolute/path/your_local_file.txt"

	err := kernel.Upload(connection, localFile, remoteFile)
	if err != nil {
		panic(err)
	}
}
