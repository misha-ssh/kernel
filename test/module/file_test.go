//go:build module

package module

import (
	"testing"

	fl "github.com/ssh-connection-manager/kernel/v2/pkg/file"

	"github.com/ssh-connection-manager/kernel/v2/test"
)

func TestCreateFile(t *testing.T) {
	testDir := test.GetDirForTests()

	randomStr := test.RandomString()

	file := fl.File{Name: randomStr, Path: testDir}
	err := file.CreateFile()
	if err != nil {
		t.Fatal("Error creating file")
	}

	if !file.IsExistFile() {
		t.Fatal("File dont created")
	}
}

func TestWriteToFile(t *testing.T) {
	testDir := test.GetDirForTests()

	randomStr := test.RandomString() + ".txt"

	file := fl.File{Name: randomStr, Path: testDir}

	err := file.CreateFile()
	if err != nil {
		t.Fatal("Error creating file")
	}

	if !file.IsExistFile() {
		t.Fatal("File dont created")
	}

	randomText := test.RandomString()
	err = file.WriteFile([]byte(randomText))
	if err != nil {
		t.Fatal("Error write to file")
	}

	fileText, err := file.ReadFile()
	if err != nil {
		t.Fatal("Error read file")
	}

	if fileText != randomText {
		t.Fatal("Error random text != text from file")
	}
}

func TestReadFile(t *testing.T) {
	files := [7]fl.File{
		{Name: test.RandomString() + ".json", Path: test.GetDirForTests()},
		{Name: test.RandomString() + ".txt", Path: test.GetDirForTests()},
		{Name: test.RandomString() + ".PNG", Path: test.GetDirForTests()},
		{Name: test.RandomString() + ".PDF", Path: test.GetDirForTests()},
		{Name: test.RandomString() + ".PDF", Path: test.GetDirForTests()},
		{Name: test.RandomString() + ".DOC", Path: test.GetDirForTests()},
		{Name: test.RandomString(), Path: test.GetDirForTests()},
	}

	for _, file := range files {
		err := file.CreateFile()
		if err != nil {
			t.Fatal("Error creating file")
		}

		if !file.IsExistFile() {
			t.Fatal("File dont created")
		}

		randomText := test.RandomString()

		err = file.WriteFile([]byte(randomText))
		if err != nil {
			t.Fatalf("Error write to file %s", file)
		}

		fileText, err := file.ReadFile()

		if err != nil {
			t.Fatal("Error read file")
		}

		if fileText != randomText {
			t.Fatalf("Error random text != text from file - is file %s", file.Path+file.Name)
		}
	}
}

func TestIsExistFile(t *testing.T) {
	testDir := test.GetDirForTests()

	randomStr := test.RandomString()
	randomStr2 := test.RandomString()

	file := fl.File{Name: randomStr, Path: testDir}
	fileWithDontExistName := fl.File{Name: randomStr2, Path: testDir}

	err := file.CreateFile()
	if err != nil {
		t.Fatal("Error creating file")
	}

	if !file.IsExistFile() {
		t.Fatal("Created file exists")
	}

	if fileWithDontExistName.IsExistFile() {
		t.Fatal("None create file is exist")
	}
}
