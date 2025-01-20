package config

import (
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"os/user"

	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/kernel/v2/pkg/file"
	"github.com/ssh-connection-manager/kernel/v2/pkg/output"
)

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		logger.Danger(err.Error())
		output.GetOutError("Error retrieving user data")
	}

	return usr.HomeDir + DirectionApp
}

func existOrCreateConfig(fl file.File) {
	err := viper.ReadInConfig()
	if err != nil {
		err := fl.CreateFile()
		if err != nil {
			errMessage := "File creation error down"

			logger.Danger(errMessage)
			output.GetOutError(errMessage)
		}

		err = viper.ReadInConfig()
		if err != nil {
			errMessage := "File creation error"

			logger.Danger(errMessage)
			output.GetOutError(errMessage)
		}
	}
}

func setConfigVariable() {
	viper.Set("NameFileConnects", NameFileConnects)
	viper.Set("NameFileCryptKey", NameFileCryptKey)
	viper.Set("NameFileLogger", NameFileLogger)
	viper.Set("FullPathConfig", getHomeDir())
	viper.Set("Separator", Separator)
	viper.Set("Space", Space)

	err := viper.WriteConfig()
	if err != nil {
		errMessage := "Error writing to configuration file"

		logger.Danger(errMessage)
		output.GetOutError(errMessage)
	}
}

func Generate() {
	viper.New()

	viper.SetConfigName(NameFileConfig)
	viper.SetConfigType(TypeFileConfig)
	viper.AddConfigPath(getHomeDir())

	err := viper.ReadInConfig()
	if err != nil {
		confPath := getHomeDir()
		confName := FullNameFileConfig

		fileConf := file.File{Path: confPath, Name: confName}

		existOrCreateConfig(fileConf)
	}

	setConfigVariable()
}
