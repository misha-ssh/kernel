package logger

import "errors"

type Logger interface {
	Error(value any)
	Debug(value any)
	Info(value any)
	Warn(value any)
}

const (
	SkipUseLevel = 1
	FileName     = "log.log"
)

var (
	ErrCreateStorage = errors.New("err at created log file")
	ErrGetOpenFile   = errors.New("err get open log file")
	ErrGetInfo       = errors.New("err get info use log")
)
