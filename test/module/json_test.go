//go:build module

package modules

import (
	fl "github.com/ssh-connection-manager/file"
	"github.com/ssh-connection-manager/file/tests"
	"github.com/ssh-connection-manager/json"
	"github.com/ssh-connection-manager/kernel/inits"
	"testing"
)

func TestWriteToJsonFile(t *testing.T) {
	inits.SetDependencies()

	path := tests.GetDirForTests()
	fileName := tests.RandomString(5) + ".json"

	file := fl.File{Name: fileName, Path: path}

	err := json.Generate(file)
	if err != nil {
		t.Fatal("Error generating json file")
	}

	connect := json.Connect{
		Alias:     tests.RandomString(5),
		Login:     tests.RandomString(5),
		Address:   tests.RandomString(5),
		Password:  tests.RandomString(5),
		CreatedAt: tests.RandomString(5),
		UpdatedAt: tests.RandomString(5),
	}

	connection := json.Connections{Connects: make([]json.Connect, 1)}

	err = connection.WriteConnectToJson(connect)
	if err != nil {
		t.Fatal("Error writing to json file " + err.Error())
	}

	err = connection.WriteConnectToJson(connect)
	if err == nil {
		t.Fatal("Error writing to json file is duplicate connect")
	}

	connectTwin := json.Connect{
		Alias:     tests.RandomString(5),
		Login:     tests.RandomString(5),
		Address:   tests.RandomString(5),
		Password:  tests.RandomString(5),
		CreatedAt: tests.RandomString(5),
		UpdatedAt: tests.RandomString(5),
	}

	err = connection.WriteConnectToJson(connectTwin)
	if err != nil {
		t.Fatal("Error writing to json file with twin connect " + err.Error())
	}
}
