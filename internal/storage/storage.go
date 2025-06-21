package storage

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

func Create(path string, filename string) error {
	file := filepath.Join(path, filename)

	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(filepath.Dir(file), os.ModePerm)
		if err != nil {
			return err
		}

		createdFile, err := os.Create(file)
		if err != nil {
			return err
		}

		defer func(createdFile *os.File) {
			err = createdFile.Close()
		}(createdFile)
	}

	return nil
}

func Delete(path string, filename string) error {
	return os.Remove(filepath.Join(path, filename))
}

func Exists(path string, filename string) bool {
	filePath := filepath.Join(path, filename)

	info, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func Get(path string, filename string) (string, error) {
	file := filepath.Join(path, filename)

	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		err = f.Close()
	}(f)

	fContent, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(fContent), nil
}

func Write(path string, filename string, data string) error {
	file := filepath.Join(path, filename)

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err = f.Close()
	}(f)

	_, err = f.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}

func GetOpenFile(path string, filename string, flags int) (*os.File, error) {
	file := filepath.Join(path, filename)

	openFile, err := os.OpenFile(file, flags, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return openFile, nil
}
