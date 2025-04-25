package logger

type Logger interface {
	Error(value any) error
	Debug(value any) error
	Info(value any) error
	Warn(value any) error
}
