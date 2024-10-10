package app

import (
	"github.com/ssh-connection-manager/json"
	"github.com/ssh-connection-manager/output"
	"github.com/ssh-connection-manager/time"
)

func Create(alias, address, login, password string) {
	var connections json.Connections

	timeNow := time.GetTime()

	connect := json.Connect{
		Alias:     alias,
		Address:   address,
		Login:     login,
		Password:  password,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	err := connections.WriteConnectToJson(connect)
	if err != nil {
		output.GetOutError(err.Error())
	}
}
