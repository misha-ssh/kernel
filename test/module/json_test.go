//go:build module

package module

import (
	"testing"

	fl "github.com/ssh-connection-manager/kernel/v2/pkg/file"

	"github.com/ssh-connection-manager/kernel/v2/inits"
	"github.com/ssh-connection-manager/kernel/v2/pkg/json"
	"github.com/ssh-connection-manager/kernel/v2/test"
)

func TestWriteToJsonFile(t *testing.T) {
	inits.SetDependencies()

	path := test.GetDirForTests()
	fileName := test.RandomString() + ".json"

	file := fl.File{Name: fileName, Path: path}

	err := json.Generate(file)
	if err != nil {
		t.Fatal("Error generating json file")
	}

	connect := json.Connect{
		Alias:     test.RandomString(),
		Login:     test.RandomString(),
		Address:   test.RandomString(),
		Password:  test.RandomString(),
		CreatedAt: test.RandomString(),
		UpdatedAt: test.RandomString(),
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
		Alias:     test.RandomString(),
		Login:     test.RandomString(),
		Address:   test.RandomString(),
		Password:  test.RandomString(),
		CreatedAt: test.RandomString(),
		UpdatedAt: test.RandomString(),
	}

	err = connection.WriteConnectToJson(connectTwin)
	if err != nil {
		t.Fatal("Error writing to json file with twin connect " + err.Error())
	}
}
