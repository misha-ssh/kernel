package app

import (
	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/pkg/json"
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
		logger.Danger(err.Error())
		output.GetOutError(err.Error())
	}
}
