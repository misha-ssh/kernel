package inits

import (
	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/crypt"
	"github.com/ssh-connection-manager/file"
	"github.com/ssh-connection-manager/json"
	"github.com/ssh-connection-manager/kernel/config"
	"github.com/ssh-connection-manager/output"
)

func generateConfigFile() {
	config.Generate()
}

func createFileConnects() {
	pathConf := viper.GetString("FullPathConfig")
	confName := viper.GetString("NameFileConnects")

	fileConnect := file.File{Path: pathConf, Name: confName}

	err := json.Generate(fileConnect)
	if err != nil {
		output.GetOutError("err create file connect")
	}
}

func generateCryptKey() {
	pathConf := viper.GetString("FullPathConfig")
	fileNameKey := viper.GetString("NameFileCryptKey")

	fileConnect := file.File{Path: pathConf, Name: fileNameKey}

	err := crypt.GenerateFileKey(fileConnect)
	if err != nil {
		output.GetOutError("err generate key")
	}
}

func SetDependencies() {
	generateConfigFile()
	createFileConnects()
	generateCryptKey()
}
