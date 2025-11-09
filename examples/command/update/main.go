package main

import (
	"time"

	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

// kernel.Update It will completely update the connection using the old alias
func main() {
	connection := &connect.Connect{
		Alias:     "test-new",
		Login:     "root",
		Password:  "password",
		Address:   "localhost",
		Type:      connect.TypeSSH,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
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
