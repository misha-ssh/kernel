package storage

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/configs/envname"
)

func TestGetAppDir(t *testing.T) {
	originalTesting := os.Getenv(envname.Testing)
	defer func() {
		err := os.Setenv(envname.Testing, originalTesting)
		if err != nil {
			t.Errorf("failed set env: %v", err)
		}
	}()

	tests := []struct {
		name         string
		want         func() string
		isSetTesting bool
	}{
		{
			name: "success - get app dir",
			want: func() string {
				usr, err := user.Current()
				if err != nil {
					t.Errorf("failed get user: %v", err)
				}

				return filepath.Join(usr.HomeDir, CharHidden+envconst.AppName)
			},
			isSetTesting: false,
		},
		{
			name: "success - get test app dir",
			want: func() string {
				return filepath.Join(os.TempDir(), CharHidden+envconst.AppName)
			},
			isSetTesting: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isSetTesting {
				err := os.Setenv(envname.Testing, "true")
				if err != nil {
					t.Errorf("failed set env: %v", err)
				}
			} else {
				err := os.Setenv(envname.Testing, "false")
				if err != nil {
					t.Errorf("failed set env: %v", err)
				}
			}

			if got := GetAppDir(); got != tt.want() {
				t.Errorf("GetAppDir() = %v, want %v", got, tt.want())
			}
		})
	}
}

func TestGetDirectionAndFilename(t *testing.T) {
	tests := []struct {
		name     string
		fullPath string
		wantDir  string
		wantFile string
	}{
		{
			name:     "simple path",
			fullPath: "/home/user/file.txt",
			wantDir:  "/home/user",
			wantFile: "file.txt",
		},
		{
			name:     "nested path",
			fullPath: "/home/user/documents/file.txt",
			wantDir:  "/home/user/documents",
			wantFile: "file.txt",
		},
		{
			name:     "current dir file",
			fullPath: "file.txt",
			wantDir:  ".",
			wantFile: "file.txt",
		},
		{
			name:     "empty path",
			fullPath: "",
			wantDir:  ".",
			wantFile: ".",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDir, gotFile := GetDirectionAndFilename(tt.fullPath)
			if gotDir != tt.wantDir {
				t.Errorf("GetDirectionAndFilename() dir = %v, want %v", gotDir, tt.wantDir)
			}
			if gotFile != tt.wantFile {
				t.Errorf("GetDirectionAndFilename() file = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}

func TestGetFullPath(t *testing.T) {
	tests := []struct {
		name      string
		direction string
		filename  string
		want      string
	}{
		{
			name:      "simple join",
			direction: "/home/user",
			filename:  "file.txt",
			want:      "/home/user/file.txt",
		},
		{
			name:      "empty dir",
			direction: "",
			filename:  "file.txt",
			want:      "file.txt",
		},
		{
			name:      "empty filename",
			direction: "/home/user",
			filename:  "",
			want:      "/home/user",
		},
		{
			name:      "both empty",
			direction: "",
			filename:  "",
			want:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFullPath(tt.direction, tt.filename); got != tt.want {
				t.Errorf("GetFullPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPrivateKeysDir(t *testing.T) {
	tests := []struct {
		name string
		want func() string
	}{
		{
			name: "success - get private keys dir",
			want: func() string {
				return filepath.Join(GetAppDir(), envconst.DirectionPrivateKeys)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPrivateKeysDir(); got != tt.want() {
				t.Errorf("GetPrivateKeysDir() = %v, want %v", got, tt.want())
			}
		})
	}
}
