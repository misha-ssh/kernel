package connect

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/ssh-connection-manager/file"
	"github.com/ssh-connection-manager/json"
)

func Ssh(c *json.Connections, alias string, fl file.File) error {
	data, err := fl.ReadFile()
	if err != nil {
		return err
	}

	err = c.SerializationJson(data)
	if err != nil {
		return err
	}

	err = c.SetDecryptData()
	if err != nil {
		return err
	}

	for _, v := range c.Connects {
		if v.Alias == alias {
			sshConnect(v.Address, v.Login, v.Password)
			return nil
		}
	}

	return errors.New("alias not found")
}

func sshConnect(address, login, password string) {
	sshCommand := "sshpass -p '" + password + "' ssh -o StrictHostKeyChecking=no -t " + login + "@" + address
	sshCmd := exec.Command("bash", "-c", sshCommand)

	sshCmd.Stdout = os.Stdout
	sshCmd.Stderr = os.Stderr
	sshCmd.Stdin = os.Stdin

	if err := sshCmd.Run(); err != nil {
		fmt.Println("Error while executing the command:", err)
	}
}
