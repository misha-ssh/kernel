//go:build module

package module

import (
	"github.com/ssh-connection-manager/kernel/v2/inits"
	fl "github.com/ssh-connection-manager/kernel/v2/pkg/file"
	"github.com/ssh-connection-manager/kernel/v2/pkg/json"
	"github.com/ssh-connection-manager/kernel/v2/test"
	"testing"
)

func TestWriteToJsonFile(t *testing.T) {
	inits.SetDependencies()

	path := test.GetDirForTests()
	fileName := test.RandomString(5) + ".json"

	file := fl.File{Name: fileName, Path: path}

	err := json.Generate(file)
	if err != nil {
		t.Fatal("Error generating json file")
	}

	connect := json.Connect{
		Alias:     test.RandomString(5),
		Login:     test.RandomString(5),
		Address:   test.RandomString(5),
		Password:  test.RandomString(5),
		CreatedAt: test.RandomString(5),
		UpdatedAt: test.RandomString(5),
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
		Alias:     test.RandomString(5),
		Login:     test.RandomString(5),
		Address:   test.RandomString(5),
		Password:  test.RandomString(5),
		CreatedAt: test.RandomString(5),
		UpdatedAt: test.RandomString(5),
	}

	err = connection.WriteConnectToJson(connectTwin)
	if err != nil {
		t.Fatal("Error writing to json file with twin connect " + err.Error())
	}
}
