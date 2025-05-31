package config

import (
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name          string
		wantErr       bool
		deleteHomeDir bool
	}{
		{
			name:          "create files with empty project dir",
			wantErr:       false,
			deleteHomeDir: true,
		},
		{
			name:          "init at created needed files",
			wantErr:       false,
			deleteHomeDir: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.deleteHomeDir {
				err := os.RemoveAll(storage.GetHomeDir())

				if err != nil {
					t.Errorf("error at delete dir = %v, wantErr %v", err, tt.wantErr)
				}
			}

			if err := Init(); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
