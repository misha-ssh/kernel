package main

import (
	"github.com/misha-ssh/kernel/pkg/ssh"
	"time"

	"github.com/misha-ssh/kernel/pkg/connect"
)

// main for success connect start make command: up-ssh
func main() {
	ssh := &ssh.Ssh{
		Connection: &connect.Connect{
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
		},
	}

	session, err := ssh.Session()
	if err != nil {
		panic(err)
	}

	err = ssh.Connect(session)
	if err != nil {
		panic(err)
	}
}
