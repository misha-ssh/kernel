package main

import (
	"github.com/ssh-connection-manager/kernel/v2/pkg/connect"
	"github.com/ssh-connection-manager/kernel/v2/pkg/kernel"
)

// kernel.Update It will completely update the connection using the old alias
func main() {
	connection := &connect.Connect{
		Alias:     "test-new",
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

	err := kernel.Update(connection, "test")
	if err != nil {
		panic(err)
	}
}
