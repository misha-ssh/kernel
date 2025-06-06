package logger

type Logger interface {
	Error(value any)
	Debug(value any)
	Info(value any)
	Warn(value any)
}

const SkipUseLevel = 3

var defaultLogger Logger

func Get() Logger {
	if defaultLogger == nil {
		defaultLogger = NewConsoleLogger()
	}

	return defaultLogger
}

func Set(logger Logger) {
	defaultLogger = logger
}

func Error(value any) {
	Get().Error(value)
}

func Debug(value any) {
	Get().Debug(value)
}

func Info(value any) {
	Get().Info(value)
}

func Warn(value any) {
	Get().Warn(value)
}
