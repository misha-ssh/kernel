package main

import "github.com/misha-ssh/kernel/pkg/connect"

// main for success connect start make command: up-ssh
// generate key ssh-keygen -b 4096 -t rsa
// ssh-copy-id root@localhost
func main() {
	connection := &connect.Connect{
		Alias:     "test",
		Login:     "root",
		Password:  "",
		Address:   "localhost",
		Type:      connect.TypeSSH,
		CreatedAt: "",
		UpdatedAt: "",
		SshOptions: &connect.SshOptions{
			Port: 22,
		},
	}

	sshConnector := &connect.Ssh{}
	session, err := sshConnector.Session(connection)
	if err != nil {
		panic(err)
	}

	err = sshConnector.Connect(session)
	if err != nil {
		panic(err)
	}
}
