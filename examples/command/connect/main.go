package main

import (
	"time"

	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

// main for success connect start make command: up-ssh
// kernel.Connect It will perform the connection by type and connection properties
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

	err := kernel.Connect(connection)
	if err != nil {
		panic(err)
	}
}
