package config

type Config interface {
	Get(key string) string
	Set(key, value string) error
	Exists(key string) bool
}
