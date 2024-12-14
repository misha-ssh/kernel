package json

import (
	"errors"

	"github.com/ssh-connection-manager/file"
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
			return err
		}

		err = CreateBaseJsonDataToFile(flConnect)
		if err != nil {
			return errors.New("err generate base json: " + err.Error())
		}
	}

	return nil
}
