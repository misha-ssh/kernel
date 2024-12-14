package json

import (
	"errors"

	"github.com/ssh-connection-manager/file"
)

func CreateBaseJsonDataToFile(fl file.File) error {
	connections := Connections{
		Connects: []Connect{},
	}

	connect, err := connections.deserializationJson()
	if err != nil {
		return errors.New("error create json: " + err.Error())
	}

	err = fl.WriteFile(connect)
	if err != nil {
		return errors.New("error write to json: " + err.Error())
	}

	return nil
}
