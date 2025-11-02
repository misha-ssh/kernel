//go:build unit

package storage

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/configs/envname"
	"github.com/stretchr/testify/require"
)

func TestGetAppDir(t *testing.T) {
	originalTesting := os.Getenv(envname.Testing)
	defer func() {
		require.NoError(t, os.Setenv(envname.Testing, originalTesting))
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
				require.NoError(t, err)

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
				require.NoError(t, os.Setenv(envname.Testing, "true"))
			} else {
				require.NoError(t, os.Setenv(envname.Testing, "false"))
			}

			require.Equal(t, GetAppDir(), tt.want())
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
			require.Equal(t, gotDir, tt.wantDir)
			require.Equal(t, gotFile, tt.wantFile)
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
			require.Equal(t, GetFullPath(tt.direction, tt.filename), tt.want)
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
			require.Equal(t, GetPrivateKeysDir(), tt.want())
		})
	}
}
