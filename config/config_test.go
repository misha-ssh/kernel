package config

import (
	"github.com/ssh-connection-manager/kernel/v2/config/envconst"
	"github.com/ssh-connection-manager/kernel/v2/internal/crypto"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"github.com/zalando/go-keyring"
	"os"
	"os/user"
	"testing"
)

func Test_Init(t *testing.T) {
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

func Test_initCryptKey(t *testing.T) {
	tests := []struct {
		name        string
		wantErr     bool
		isDeleteKey bool
	}{
		{
			name:        "get save crypt key",
			wantErr:     false,
			isDeleteKey: false,
		},
		{
			name:        "empty key - set crypt key",
			wantErr:     false,
			isDeleteKey: true,
		},
		{
			name:        "get generated crypt key",
			wantErr:     false,
			isDeleteKey: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentUser, err := user.Current()
			if err != nil {
				t.Errorf("error get user = %v, wantErr %v", err, tt.wantErr)
			}

			username := currentUser.Username

			if tt.isDeleteKey {
				err = keyring.Delete(envconst.NameServiceCryptKey, username)
				if err != nil {
					t.Errorf("error at delete key = %v, wantErr %v", err, tt.wantErr)
				}
			}

			if err = initCryptKey(); (err != nil) != tt.wantErr {
				t.Errorf("initCryptKey() error = %v, wantErr %v", err, tt.wantErr)
			}

			key, err := keyring.Get(envconst.NameServiceCryptKey, username)
			if (err != nil) != tt.wantErr {
				t.Errorf("error get key = %v, wantErr %v", err, tt.wantErr)
			}

			if len(key) != crypto.SizeKey {
				t.Errorf("invalid key = %v, want %v", key, crypto.SizeKey)
			}
		})
	}
}

func Test_initFileConfig(t *testing.T) {
	tests := []struct {
		name           string
		wantErr        bool
		isDeleteConfig bool
	}{
		{
			name:           "config is exist",
			wantErr:        false,
			isDeleteConfig: false,
		},
		{
			name:           "config is empty - create config",
			wantErr:        false,
			isDeleteConfig: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileStorage := &storage.FileStorage{
				Direction: storage.GetHomeDir(),
			}

			if tt.isDeleteConfig {
				err := fileStorage.Delete(envconst.FilenameConfig)
				if err != nil {
					t.Errorf("error at delete config = %v, wantErr %v", err, tt.wantErr)
				}
			}

			if err := initFileConfig(); (err != nil) != tt.wantErr {
				t.Errorf("initFileConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !fileStorage.Exists(envconst.FilenameConfig) {
				t.Errorf("config file not exist")
			}
		})
	}
}

func Test_initFileConnections(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := initFileConnections(); (err != nil) != tt.wantErr {
				t.Errorf("initFileConnections() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
