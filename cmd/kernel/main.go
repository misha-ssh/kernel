package main

import "github.com/ssh-connection-manager/kernel/v2/internal/connect"

func main() {
	conn := connect.SshConnect{}
	conn.Connect()
}
