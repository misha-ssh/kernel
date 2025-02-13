package storage

type Storage interface {
	Exists(filename string) bool
	Create(filename string) error
	Get(filename string) (string, error)
	Delete(filename string) error
	Write(filename string, data []byte) error
}
