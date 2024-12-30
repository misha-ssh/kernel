package logger

import "github.com/ssh-connection-manager/kernel/v2/pkg/file"

var fl file.File

func SetFile(file file.File) {
	fl = file
}

func GetFile() file.File {
	return fl
}
