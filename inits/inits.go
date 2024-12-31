package inits

import (
	"github.com/spf13/viper"

	"github.com/ssh-connection-manager/kernel/v2/config"
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/pkg/crypt"
	"github.com/ssh-connection-manager/kernel/v2/pkg/file"
	"github.com/ssh-connection-manager/kernel/v2/pkg/json"
	"github.com/ssh-connection-manager/kernel/v2/pkg/output"
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

func generateLogFile() {
	pathConf := viper.GetString("FullPathConfig")
	fileNameKey := viper.GetString("NameFileLogger")

	fileLogger := file.File{Path: pathConf, Name: fileNameKey}

	err := logger.GenerateFile(fileLogger)
	if err != nil {
		output.GetOutError("err generate logger")
	}
}

func SetDependencies() {
	generateConfigFile()
	createFileConnects()
	generateCryptKey()
	generateLogFile()
}
