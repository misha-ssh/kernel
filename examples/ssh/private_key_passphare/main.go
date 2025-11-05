package main

import "github.com/misha-ssh/kernel/pkg/connect"

//todo add make command for passphare
//todo update this example

// main for success connect start make command: up-ssh
// generate key ssh-keygen -b 4096 -t rsa
// ssh-copy-id root@localhost
func main() {
	ssh := &connect.Ssh{
		Connection: &connect.Connect{
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
