package main

import (
	"time"

	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

// kernel.Create It will record the connection in a file with connections
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

	err := kernel.Create(connection)
	if err != nil {
		panic(err)
	}
}
