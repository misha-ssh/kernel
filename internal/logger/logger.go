package logger

import (
	"sync"

	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

type Logger interface {
	Error(value any)
	Debug(value any)
	Info(value any)
	Warn(value any)
}

// todo убрать эту константу и вызывать логику без нее
const SkipUseLevel = 1

type TypeLogger int

const (
	StorageLoggerType TypeLogger = iota
	ConsoleLoggerType
	CombinedLoggerType
)

var (
	defaultLogger Logger
	once          sync.Once
)

func Init(loggerType TypeLogger, storage storage.Storage) {
	once.Do(func() {
		switch loggerType {
		case StorageLoggerType:
			defaultLogger = NewStorageLogger(storage)
		case ConsoleLoggerType:
			defaultLogger = NewConsoleLogger()
		case CombinedLoggerType:
			defaultLogger = NewCombinedLogger(
				NewStorageLogger(storage),
				NewConsoleLogger(),
			)
		default:
			defaultLogger = NewConsoleLogger()
		}
	})
}

func Get() Logger {
	if defaultLogger == nil {
		defaultLogger = NewConsoleLogger()
	}

	return defaultLogger
}

func SetLogger(logger Logger) {
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
