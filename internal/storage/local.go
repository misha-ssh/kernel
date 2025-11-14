package storage

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Local struct{}

var ErrEmptyDirectory = errors.New("empty directory")
var ErrDeleteDirectory = errors.New("get dir, delete only file")

func NewLocal() *Local {
	return new(Local)
}

// Create creates a new file at the specified path, including parent directories if needed.
// Returns error if file creation fails.
func (l *Local) Create(path string, filename string) error {
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
		return l.Write(path, filename, "")
	}

	return nil
}

// Delete removes the specified file. Returns error if deletion fails.
func (l *Local) Delete(path string, filename string) error {
	file := filepath.Join(path, filename)
	info, err := os.Stat(file)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return ErrDeleteDirectory
	}

	return os.Remove(filepath.Join(path, filename))
}

// Exists checks if a file exists at the given path and is not a directory.
// Returns boolean indicating existence.
func (l *Local) Exists(path string, filename string) bool {
	info, err := os.Stat(filepath.Join(path, filename))
	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return !info.IsDir()
}

// Get reads and returns the contents of a file as a string.
// Returns error if file cannot be read.
func (l *Local) Get(path string, filename string) (string, error) {
	data, err := os.ReadFile(filepath.Join(path, filename))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Write saves data to a file, overwriting existing content.
// Creates file if it doesn't exist. Returns error on failure.
func (l *Local) Write(path string, filename string, data string) error {
	err := os.WriteFile(filepath.Join(path, filename), []byte(data), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// GetOpenFile opens a file with specified flags (os.O_RDWR, etc.) and returns the file handle.
// Returns error if file cannot be opened.
func (l *Local) GetOpenFile(path string, filename string, flags int) (*os.File, error) {
	file := filepath.Join(path, filename)

	openFile, err := os.OpenFile(file, flags, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return openFile, nil
}
