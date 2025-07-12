package main

import (
	"fmt"

	"github.com/ssh-connection-manager/kernel/pkg/kernel"
)

// kernel.List outputs the connection from the file
func main() {
	connections, err := kernel.List()
	if err != nil {
		panic(err)
	}

	fmt.Println(connections)
}
