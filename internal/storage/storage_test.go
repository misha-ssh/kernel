package storage

import (
	"os"
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
				if (err != nil) != tt.wantErr {
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
				if (err != nil) != tt.wantErr {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				}

				err = Write(tt.args.direction, tt.args.filename, tt.want)
				if (err != nil) != tt.wantErr {
					t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				}
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
