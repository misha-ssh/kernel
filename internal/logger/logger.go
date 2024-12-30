package logger

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ssh-connection-manager/kernel/v2/pkg/file"
	"github.com/ssh-connection-manager/kernel/v2/pkg/output"
)

func GenerateFile(fl file.File) error {
	SetFile(fl)
	logFile := GetFile()

	if !logFile.IsExistFile() {
		err := logFile.CreateFile()
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO переписать логику на пакет file
func getOpenLogFile(fl file.File) (*os.File, error) {
	path := filepath.Join(fl.Path, fl.Name)

	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

// Info Danger TODO написать одну логику а функции только различаются строкой
func Info(message string) {
	logFile := GetFile()
	openLogFile, err := getOpenLogFile(logFile)
	if err != nil {
		output.GetOutError("dont open log file")
	}

	infoLog := log.New(openLogFile, "[info] ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	infoLog.Println(message)
}

func Danger(message string) {
	logFile := GetFile()
	openLogFile, err := getOpenLogFile(logFile)
	if err != nil {
		output.GetOutError("dont open log file")
	}

	errorLog := log.New(openLogFile, "[error] ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	errorLog.Println(message)
}
