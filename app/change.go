package app

import (
	"github.com/ssh-connection-manager/json"
	"github.com/ssh-connection-manager/output"
	"github.com/ssh-connection-manager/time"
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
		output.GetOutError("err update")
	}
}
