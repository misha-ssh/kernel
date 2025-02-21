package app

import (
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/pkg/json"
)

func Delete(alias string) {
	var connects json.Connections

	err := connects.DeleteConnectToJson(alias)
	if err != nil {
		logger.Danger(err.Error())
		output.GetOutError(err.Error())
	}
}
