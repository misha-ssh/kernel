package storage

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/config/envconst"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	type args struct {
		direction string
		filename  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create file with valida data",
			args: args{
				direction: t.TempDir(),
				filename:  "test.txt",
			},
			wantErr: false,
		},
		{
			name: "create file with empty filename - create dir",
			args: args{
				direction: t.TempDir(),
				filename:  "",
			},
			wantErr: false,
		},
		{
			name: "create file with empty two args",
			args: args{
				direction: "",
				filename:  "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Create(tt.args.direction, tt.args.filename)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		direction string
		filename  string
	}
	tests := []struct {
		name         string
		args         args
		isCreateFile bool
		wantErr      bool
	}{
		{
			name: "delete existing file",
			args: args{
				direction: t.TempDir(),
				filename:  "test.txt",
			},
			isCreateFile: true,
			wantErr:      false,
		},
		{
			name: "delete dont created file",
			args: args{
				direction: t.TempDir(),
				filename:  "nonexistent.txt",
			},
			isCreateFile: false,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isCreateFile {
				err := Create(tt.args.direction, tt.args.filename)
				assert.NoError(t, err)
			}

			err := Delete(tt.args.direction, tt.args.filename)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestExists(t *testing.T) {
	type args struct {
		direction string
		filename  string
	}
	tests := []struct {
		name         string
		args         args
		isCreateFile bool
		want         bool
	}{
		{
			name: "is exist created file",
			args: args{
				direction: t.TempDir(),
				filename:  "test.txt",
			},
			isCreateFile: true,
			want:         true,
		},
		{
			name: "is exist dont created file",
			args: args{
				direction: t.TempDir(),
				filename:  "nonexistent.txt",
			},
			isCreateFile: false,
			want:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isCreateFile {
				err := Create(tt.args.direction, tt.args.filename)
				assert.Nil(t, err)
			}

			got := Exists(tt.args.direction, tt.args.filename)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		direction string
		filename  string
	}
	tests := []struct {
		name         string
		args         args
		want         string
		isCreateFile bool
		wantErr      bool
	}{
		{
			name: "get dont empty file",
			args: args{
				direction: t.TempDir(),
				filename:  "test.txt",
			},
			want:         "test data",
			isCreateFile: true,
			wantErr:      false,
		},
		{
			name: "get empty file",
			args: args{
				direction: t.TempDir(),
				filename:  "nonexistent.txt",
			},
			want:         "",
			isCreateFile: true,
			wantErr:      false,
		},
		{
			name: "get dont created file",
			args: args{
				direction: t.TempDir(),
				filename:  "nonexistent.txt",
			},
			want:         "",
			isCreateFile: false,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isCreateFile {
				err := Create(tt.args.direction, tt.args.filename)
				assert.NoError(t, err)

				err = Write(tt.args.direction, tt.args.filename, tt.want)
				assert.NoError(t, err)
			}

			got, err := Get(tt.args.direction, tt.args.filename)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestWrite(t *testing.T) {
	type args struct {
		direction string
		filename  string
		data      string
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantContent string
	}{
		{
			name: "write to new file",
			args: args{
				direction: t.TempDir(),
				filename:  "test.txt",
				data:      "Hello, World!",
			},
			wantErr:     false,
			wantContent: "Hello, World!",
		},
		{
			name: "write to existing file (overwrite)",
			args: args{
				direction: t.TempDir(),
				filename:  "test.txt",
				data:      "New content",
			},
			wantErr:     false,
			wantContent: "New content",
		},
		{
			name: "write empty data to new file",
			args: args{
				direction: t.TempDir(),
				filename:  "empty.txt",
				data:      "",
			},
			wantErr:     false,
			wantContent: "",
		},
		{
			name: "write to invalid filename",
			args: args{
				direction: t.TempDir(),
				filename:  "",
				data:      "Invalid",
			},
			wantErr:     true,
			wantContent: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Write(tt.args.direction, tt.args.filename, tt.args.data)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				content, err := Get(tt.args.direction, tt.args.filename)
				assert.NoError(t, err)

				assert.Equal(t, tt.wantContent, content)
			}
		})
	}
}

func TestGetOpenFile(t *testing.T) {
	type args struct {
		direction string
		filename  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "write to new file",
			args: args{
				direction: t.TempDir(),
				filename:  "test.txt",
			},
			wantErr: false,
		},
		{
			name: "error on invalid directory",
			args: args{
				direction: "invalidDir" + t.TempDir(),
				filename:  "test.txt",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flags := os.O_WRONLY | os.O_APPEND | os.O_CREATE
			got, err := GetOpenFile(tt.args.direction, tt.args.filename, flags)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)

				_, err = got.Write([]byte("test"))
				assert.NoError(t, err)

				err = got.Close()
				assert.NoError(t, err)

				fileIsExists := Exists(tt.args.direction, tt.args.filename)
				assert.True(t, fileIsExists)
			}
		})
	}
}

func TestGetAppDir(t *testing.T) {
	tests := []struct {
		name string
		want func() string
	}{
		{
			name: "get success app dir",
			want: func() string {
				usr, err := user.Current()
				if err != nil {
					panic(err)
				}

				hiddenDir := "." + envconst.AppName

				return filepath.Join(usr.HomeDir, hiddenDir)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want()
			assert.Equalf(t, want, GetAppDir(), "GetAppDir()")
		})
	}
}
