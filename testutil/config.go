package testutil

import (
	"os"
	"path/filepath"
)

func CreateSSHConfig(tmpDir string, file string) error {
	sourceData, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	destPath := filepath.Join(tmpDir, "config")
	if err = os.WriteFile(destPath, sourceData, 0644); err != nil {
		return err
	}

	return nil
}
