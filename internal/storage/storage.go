package storage

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var ErrEmptyDirectory = errors.New("empty directory")

// Create creates a new file at the specified path, including parent directories if needed.
// Returns error if file creation fails.
func Create(path string, filename string) error {
	if strings.TrimSpace(path) == "" {
		return ErrEmptyDirectory
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if strings.TrimSpace(filename) != "" {
		file := filepath.Join(path, filename)

		f, err := os.Create(file)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	return nil
}

// Delete removes the specified file. Returns error if deletion fails.
func Delete(path string, filename string) error {
	return os.Remove(filepath.Join(path, filename))
}

// Exists checks if a file exists at the given path and is not a directory.
// Returns boolean indicating existence.
func Exists(path string, filename string) bool {
	filePath := filepath.Join(path, filename)

	info, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// Get reads and returns the contents of a file as a string.
// Returns error if file cannot be read.
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

// Write saves data to a file, overwriting existing content.
// Creates file if it doesn't exist. Returns error on failure.
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

// GetOpenFile opens a file with specified flags (os.O_RDWR, etc.) and returns the file handle.
// Returns error if file cannot be opened.
func GetOpenFile(path string, filename string, flags int) (*os.File, error) {
	file := filepath.Join(path, filename)

	openFile, err := os.OpenFile(file, flags, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return openFile, nil
}
