package app

import (
	"github.com/ssh-connection-manager/kernel/pkg/json"
	"github.com/ssh-connection-manager/kernel/pkg/output"
)

func Delete(alias string) {
	var connects json.Connections

	err := connects.DeleteConnectToJson(alias)
	if err != nil {
		output.GetOutError(err.Error())
	}
}
