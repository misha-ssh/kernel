package logger

type Status string

const (
	ErrorStatus Status = "ERROR"
	DebugStatus Status = "DEBUG"
	InfoStatus  Status = "INFO"
	WarnStatus  Status = "WARN"
)
