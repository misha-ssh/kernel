package main

import "github.com/ssh-connection-manager/kernel/v2/internal/connect"

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

	sshConnect := connect.NewSshConnect()
	session, err := sshConnect.Connect(connection)
	if err != nil {
		panic(err)
	}

	err = session.Shell()
	if err != nil {
		panic(err)
	}

	err = session.Wait()
	if err != nil {
		panic(err)
	}

	defer session.Close()
}
