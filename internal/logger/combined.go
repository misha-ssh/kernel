package logger

type CombinedLogger struct {
	loggers []Logger
}

func NewCombinedLogger(loggers ...Logger) *CombinedLogger {
	return &CombinedLogger{
		loggers: loggers,
	}
}

func (cl *CombinedLogger) Error(value any) {
	for _, logger := range cl.loggers {
		logger.Error(value)
	}
}

func (cl *CombinedLogger) Debug(value any) {
	for _, logger := range cl.loggers {
		logger.Debug(value)
	}
}

func (cl *CombinedLogger) Info(value any) {
	for _, logger := range cl.loggers {
		logger.Info(value)
	}
}

func (cl *CombinedLogger) Warn(value any) {
	for _, logger := range cl.loggers {
		logger.Warn(value)
	}
}
