package config

type config interface {
	Get(key string) string
	Set(key, value string)
	Exists(key string) bool
}
