package main

import "github.com/misha-ssh/kernel/pkg/connect"

// main for success connect start make command: up-ssh-port
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
			Port:       2222,
			PrivateKey: "",
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
