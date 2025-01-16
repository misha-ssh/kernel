//go:build module

package module

import (
	"testing"

	fl "github.com/ssh-connection-manager/kernel/v2/pkg/file"

	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
	"github.com/ssh-connection-manager/kernel/v2/test"
)

func TestGenerateLogFile(t *testing.T) {
	filePath := test.GetDirForTests()
	fileName := test.RandomString() + ".log"

	logFile := fl.File{
		Name: fileName,
		Path: filePath,
	}

	err := logger.GenerateFile(logFile)
	if err != nil {
		t.Fatal("Error creating file for logs")
	}

	if !logFile.IsExistFile() {
		t.Fatal("Log file dont created")
	}
}

func TestDanger(t *testing.T) {
	filePath := test.GetDirForTests()
	fileName := test.RandomString() + ".log"

	logFile := fl.File{
		Name: fileName,
		Path: filePath,
	}

	err := logger.GenerateFile(logFile)
	if err != nil {
		t.Fatal("Error creating file for logs")
	}

	if !logFile.IsExistFile() {
		t.Fatal("Log file dont created")
	}

	textMessage := test.RandomString()
	logger.Danger(textMessage)
}
