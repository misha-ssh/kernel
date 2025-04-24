package logger

type Logger interface {
	log(value any) error
}
