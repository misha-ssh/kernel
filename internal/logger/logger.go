package logger

// Logger interface defines standard logging methods
type Logger interface {
	Error(value any) // Logs error messages
	Debug(value any) // Logs debug information
	Info(value any)  // Logs general information
	Warn(value any)  // Logs warning messages
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
