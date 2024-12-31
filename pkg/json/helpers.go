package json

import (
	"errors"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"

	"github.com/ssh-connection-manager/kernel/v2/pkg/file"
)

func CreateBaseJsonDataToFile(fl file.File) error {
	connections := Connections{
		Connects: []Connect{},
	}

	connect, err := connections.deserializationJson()
	if err != nil {
		logger.Danger(err.Error())
		return errors.New("error create json: " + err.Error())
	}

	err = fl.WriteFile(connect)
	if err != nil {
		logger.Danger(err.Error())
		return errors.New("error write to json: " + err.Error())
	}

	return nil
}
