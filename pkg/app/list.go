package app

import (
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/pkg/json"
)

func List() [][]string {
	var connections json.Connections

	connectionsData, err := connections.GetDataForListConnect()
	if err != nil {
		logger.Danger(err.Error())
		output.GetOutError(err.Error())
	}

	return connectionsData
}
