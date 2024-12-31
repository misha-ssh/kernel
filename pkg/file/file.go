package file

import (
	"errors"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"io"
	"os"
	"path/filepath"
)

type File struct {
	Path string
	Name string
}

func (fl *File) IsExistFile() bool {
	_, err := fl.ReadFile()
	return err == nil
}

func (fl *File) fullPath() string {
	return filepath.Join(fl.Path, fl.Name)
}

func (fl *File) CreateFile() error {
	file := fl.fullPath()

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(filepath.Dir(file), os.ModePerm)
		if err != nil {
			logger.Danger(err.Error())
			return err
		}

		createdFile, err := os.Create(file)
		if err != nil {
			logger.Danger(err.Error())
			return err
		}

		defer func(createdFile *os.File) {
			err = createdFile.Close()
		}(createdFile)
	}

	return nil
}

func (fl *File) ReadFile() (string, error) {
	file := fl.fullPath()

	f, err := os.Open(file)
	if err != nil {
		logger.Danger(err.Error())
		return "", err
	}
	defer func(f *os.File) {
		err = f.Close()
	}(f)

	fContent, err := io.ReadAll(f)
	if err != nil {
		logger.Danger(err.Error())
		return "", err
	}

	return string(fContent), nil
}

func (fl *File) WriteFile(rowData []byte) error {
	file := fl.fullPath()

	f, err := os.OpenFile(file, os.O_WRONLY, 0666)
	if err != nil {
		logger.Danger(err.Error())
		return err
	}
	defer func(f *os.File) {
		err = f.Close()
	}(f)

	_, err = f.Write(rowData)
	if err != nil {
		logger.Danger(err.Error())
		return err
	}

	return nil
}
