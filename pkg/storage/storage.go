package storage

import "os"

type Storage interface {
	Exists(filename string) bool
	Create(filename string) error
	Get(filename string) (string, error)
	Delete(filename string) error
	Write(filename string, data string) error
	GetOpenFile(filename string) (*os.File, error)
	WriteToOpenFile(openFile *os.File, data string) error
}
