package storage

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	Direction string
}

func (s *LocalStorage) Create(filename string) error {
	file := filepath.Join(s.Direction, filename)

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

func (s *LocalStorage) Delete(filename string) error {
	return os.Remove(filepath.Join(s.Direction, filename))
}

func (s *LocalStorage) Exists(filename string) bool {
	filePath := filepath.Join(s.Direction, filename)

	info, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func (s *LocalStorage) Get(filename string) (string, error) {
	file := filepath.Join(s.Direction, filename)

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

func (s *LocalStorage) Write(filename string, data string) error {
	file := filepath.Join(s.Direction, filename)

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

func (s *LocalStorage) GetOpenFile(filename string, flags int) (*os.File, error) {
	file := filepath.Join(s.Direction, filename)

	openFile, err := os.OpenFile(file, flags, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return openFile, nil
}
