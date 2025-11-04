package main

import "github.com/misha-ssh/kernel/pkg/connect"

// main for success connect start make command: up-ssh-port
func main() {
	ssh := &connect.Ssh{
		Connection: &connect.Connect{
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
