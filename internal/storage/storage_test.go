package storage

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
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
			name: "success - create file",
			args: args{
				direction: t.TempDir(),
				filename:  "test.txt",
			},
			wantErr: false,
		},
		{
			name: "success - create dir",
			args: args{
				direction: t.TempDir() + "/new_dir",
				filename:  "",
			},
			wantErr: false,
		},
		{
			name: "fail - empty dir",
			args: args{
				direction: "",
				filename:  "new.txt",
			},
			wantErr: true,
		},
		{
			name: "fail - empty dir and filename",
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

			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
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
			name: "success - delete file",
			args: args{
				direction: t.TempDir(),
				filename:  "test.txt",
			},
			isCreateFile: true,
			wantErr:      false,
		},
		{
			name: "fail - delete non exists file",
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
				if err != nil {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			err := Delete(tt.args.direction, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
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
				if err != nil {
					t.Errorf("Create() error = %v", err)
				}
			}

			got := Exists(tt.args.direction, tt.args.filename)
			if got != tt.want {
				t.Errorf("got: %v != want: %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tempDir := t.TempDir()

	testFiles := map[string][]byte{
		"test.txt":         []byte("test data"),
		"empty.txt":        []byte(""),
		"large.txt":        make([]byte, 1024*1024),
		"large-repeat.txt": bytes.Repeat([]byte("x"), 1024*1024),
	}

	for filename, data := range testFiles {
		filePath := filepath.Join(tempDir, filename)
		err := os.WriteFile(filePath, data, 0644)
		if err != nil {
			t.Fatalf("WriteFile() %s: error: %v", filename, err)
		}
	}

	tests := []struct {
		name      string
		direction string
		filename  string
		want      string
		wantErr   bool
	}{
		{
			name:      "success - read test file",
			direction: tempDir,
			filename:  "test.txt",
			want:      string(testFiles["test.txt"]),
			wantErr:   false,
		},
		{
			name:      "success - read empty file",
			direction: tempDir,
			filename:  "empty.txt",
			want:      string(testFiles["empty.txt"]),
			wantErr:   false,
		},
		{
			name:      "success - read large file",
			direction: tempDir,
			filename:  "large.txt",
			want:      string(testFiles["large.txt"]),
			wantErr:   false,
		},
		{
			name:      "success - read large file",
			direction: tempDir,
			filename:  "large-repeat.txt",
			want:      string(testFiles["large-repeat.txt"]),
			wantErr:   false,
		},
		{
			name:      "fail - non-existent file",
			direction: tempDir,
			filename:  "nonexistent.txt",
			want:      "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.direction, tt.filename)

			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("got: %v != want: %v", got, tt.want)
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
		name    string
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "write to new file",
			args: args{
				direction: t.TempDir(),
				filename:  "test.txt",
				data:      "Hello, World!",
			},
			wantErr: false,
			want:    "Hello, World!",
		},
		{
			name: "write to existing file (overwrite)",
			args: args{
				direction: t.TempDir(),
				filename:  "test.txt",
				data:      "New content",
			},
			wantErr: false,
			want:    "New content",
		},
		{
			name: "write empty data to new file",
			args: args{
				direction: t.TempDir(),
				filename:  "empty.txt",
				data:      "",
			},
			wantErr: false,
			want:    "",
		},
		{
			name: "write to invalid filename",
			args: args{
				direction: t.TempDir(),
				filename:  "",
				data:      "Invalid",
			},
			wantErr: true,
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Write(tt.args.direction, tt.args.filename, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := Get(tt.args.direction, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("got: %v != want: %v", got, tt.want)
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
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOpenFile() error = %v, wantErr %v", err, tt.wantErr)
			}

			_, err = got.Write([]byte("test"))
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			err = got.Close()
			if (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}

			fileIsExists := Exists(tt.args.direction, tt.args.filename)

			if !fileIsExists != tt.wantErr {
				t.Errorf("file not exists")
			}
		})
	}
}
