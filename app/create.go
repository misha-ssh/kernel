package app

import (
	"github.com/ssh-connection-manager/kernel/pkg/json"
	"github.com/ssh-connection-manager/kernel/pkg/output"
	"github.com/ssh-connection-manager/kernel/pkg/time"
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
