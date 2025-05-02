package config

//TODO добавить load для установки значений если их нет
type Config interface {
	Get(key string) string
	Set(key, value string) error
	Exists(key string) bool
}
