package logger

type StatusLog string

const (
	ErrorStatus StatusLog = "ERROR"
	DebugStatus StatusLog = "DEBUG"
	InfoStatus  StatusLog = "INFO"
	WarnStatus  StatusLog = "WARN"
)
