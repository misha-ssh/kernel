package app

import (
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/internal/time"
	"github.com/ssh-connection-manager/kernel/v2/pkg/json"
	"github.com/ssh-connection-manager/kernel/v2/pkg/output"
)

func Change(
	oldAlias string,
	alias, address, login, password string) {
	var connections json.Connections

	timeNow := time.GetTime()

	connect := json.Connect{
		Alias:     alias,
		Address:   address,
		Login:     login,
		Password:  password,
		UpdatedAt: timeNow,
	}

	err := connections.UpdateConnectJson(oldAlias, connect)
	if err != nil {
		logger.Danger(err.Error())
		output.GetOutError(err.Error())
	}
}
