package app

import (
	"github.com/spf13/viper"
	"github.com/ssh-connection-manager/kernel/v2/pkg/connect"
	"github.com/ssh-connection-manager/kernel/v2/pkg/file"
	"github.com/ssh-connection-manager/kernel/v2/pkg/json"
	"github.com/ssh-connection-manager/kernel/v2/pkg/output"
)

func Connect(alias string) {
	var connections json.Connections

	filePath := viper.GetString("FullPathConfig")
	fileName := viper.GetString("NameFileConnects")

	fileConnect := file.File{Path: filePath, Name: fileName}

	err := connect.Ssh(&connections, alias, fileConnect)
	if err != nil {
		output.GetOutError("err connect")
	}
}
