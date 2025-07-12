package main

import (
	"github.com/ssh-connection-manager/kernel/pkg/connect"
	"github.com/ssh-connection-manager/kernel/pkg/kernel"
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
		CreatedAt: "",
		UpdatedAt: "",
		SshOptions: &connect.SshOptions{
			Port:       22,
			PrivateKey: "",
		},
	}

	err := kernel.Connect(connection)
	if err != nil {
		panic(err)
	}
}
