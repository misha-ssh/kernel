package main

import "github.com/ssh-connection-manager/kernel/v2/internal/connect"

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
			PrivateKey: false,
		},
	}
	conn := connect.SshConnect{}

	err := conn.Connect(connection)
	if err != nil {
		panic(err)
	}
}
