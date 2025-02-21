package json

import (
	"errors"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
)

var fl file.File

func SetFile(file file.File) {
	fl = file
}

func GetFile() file.File {
	return fl
}

func Generate(fl file.File) error {
	SetFile(fl)
	flConnect := GetFile()

	if !flConnect.IsExistFile() {
		err := flConnect.CreateFile()
		if err != nil {
			logger.Danger(err.Error())
			return err
		}

		err = CreateBaseJsonDataToFile(flConnect)
		if err != nil {
			logger.Danger(err.Error())
			return errors.New("err generate base json: " + err.Error())
		}
	}

	return nil
}
