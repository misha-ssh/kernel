package crypt

var fl file.File

func SetFile(file file.File) {
	fl = file
}

func GetFile() file.File {
	return fl
}
