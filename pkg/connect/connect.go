package connect

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/pkg/file"
	"github.com/ssh-connection-manager/kernel/v2/pkg/json"
)

func Ssh(c *json.Connections, alias string, fl file.File) error {
	data, err := fl.ReadFile()
	if err != nil {
		logger.Danger(err.Error())
		return err
	}

	err = c.SerializationJson(data)
	if err != nil {
		logger.Danger(err.Error())
		return err
	}

	err = c.SetDecryptData()
	if err != nil {
		logger.Danger(err.Error())
		return err
	}

	for _, v := range c.Connects {
		if v.Alias == alias {
			sshConnect(v.Address, v.Login, v.Password)
			return nil
		}
	}

	errText := "alias not found"

	logger.Danger(errText)
	return errors.New(errText)
}

func sshConnect(address, login, password string) {
	sshCommand := "sshpass -p '" + password + "' ssh -o StrictHostKeyChecking=no -t " + login + "@" + address
	sshCmd := exec.Command("bash", "-c", sshCommand)

	sshCmd.Stdout = os.Stdout
	sshCmd.Stderr = os.Stderr
	sshCmd.Stdin = os.Stdin

	if err := sshCmd.Run(); err != nil {
		logger.Danger(err.Error())
		fmt.Println("Error while executing the command:", err)
	}
}
