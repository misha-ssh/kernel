package main

import (
	"github.com/ssh-connection-manager/kernel/v2/pkg/connect"
	"github.com/ssh-connection-manager/kernel/v2/pkg/kernel"
)

// kernel.Create It will record the connection in a file with connections
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

	err := kernel.Create(connection)
	if err != nil {
		panic(err)
	}
}
