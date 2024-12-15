package app

import (
	"github.com/ssh-connection-manager/kernel/pkg/json"
	"github.com/ssh-connection-manager/kernel/pkg/output"
)

func List() [][]string {
	var connections json.Connections

	connectionsData, err := connections.GetDataForListConnect()
	if err != nil {
		output.GetOutError(err.Error())
	}

	return connectionsData
}
