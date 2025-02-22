package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalStorage_Create(t *testing.T) {
	type args struct {
		filename  string
		direction string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create file with valida data",
			args: args{
				filename:  "test.txt",
				direction: t.TempDir(),
			},
			wantErr: false,
		},
		{
			name: "create file with empty filename - create dir",
			args: args{
				filename:  "",
				direction: t.TempDir(),
			},
			wantErr: false,
		},
		{
			name: "create file with empty two args",
			args: args{
				filename:  "",
				direction: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{}

			err := s.Create(
				tt.args.filename,
				tt.args.direction,
			)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLocalStorage_Delete(t *testing.T) {
	type args struct {
		filename  string
		direction string
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
				filename:  "test.txt",
				direction: t.TempDir(),
			},
			isCreateFile: true,
			wantErr:      false,
		},
		{
			name: "delete dont created file",
			args: args{
				filename:  "nonexistent.txt",
				direction: t.TempDir(),
			},
			isCreateFile: false,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{}

			if tt.isCreateFile {
				err := s.Create(tt.args.filename, tt.args.direction)
				assert.NoError(t, err)
			}

			err := s.Delete(tt.args.filename, tt.args.direction)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLocalStorage_Exists(t *testing.T) {
	type args struct {
		filename  string
		direction string
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
				filename:  "test.txt",
				direction: t.TempDir(),
			},
			isCreateFile: true,
			want:         true,
		},
		{
			name: "is exist dont created file",
			args: args{
				filename:  "nonexistent.txt",
				direction: t.TempDir(),
			},
			isCreateFile: false,
			want:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{}

			if tt.isCreateFile {
				err := s.Create(tt.args.filename, tt.args.direction)
				assert.Nil(t, err)
			}

			got := s.Exists(tt.args.filename, tt.args.direction)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLocalStorage_Get(t *testing.T) {
	type args struct {
		filename  string
		direction string
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
				filename:  "test.txt",
				direction: t.TempDir(),
			},
			want:         "test data",
			isCreateFile: true,
			wantErr:      false,
		},
		{
			name: "get empty file",
			args: args{
				filename:  "nonexistent.txt",
				direction: t.TempDir(),
			},
			want:         "",
			isCreateFile: true,
			wantErr:      false,
		},
		{
			name: "get dont created file",
			args: args{
				filename:  "nonexistent.txt",
				direction: t.TempDir(),
			},
			want:         "",
			isCreateFile: false,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{}

			if tt.isCreateFile {
				err := s.Create(tt.args.filename, tt.args.direction)
				assert.NoError(t, err)

				err = s.Write(tt.args.filename, tt.args.direction, tt.want)
				assert.NoError(t, err)
			}

			got, err := s.Get(tt.args.filename, tt.args.direction)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestLocalStorage_Write(t *testing.T) {
	type args struct {
		filename  string
		direction string
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
				filename:  "test.txt",
				direction: t.TempDir(),
				data:      "test",
			},
			wantErr:     false,
			wantContent: "test",
		},
		{
			name: "write to existing file (overwrite)",
			args: args{
				filename:  "test.txt",
				direction: t.TempDir(),
				data:      "test",
			},
			wantErr:     false,
			wantContent: "test",
		},
		{
			name: "write empty data to new file",
			args: args{
				filename:  "test.txt",
				direction: t.TempDir(),
				data:      "",
			},
			wantErr:     false,
			wantContent: "",
		},
		{
			name: "write to invalid filename",
			args: args{
				filename:  "",
				direction: t.TempDir(),
				data:      "Hello, World!",
			},
			wantErr:     true,
			wantContent: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{}

			err := s.Write(tt.args.filename, tt.args.direction, tt.args.data)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				content, err := s.Get(tt.args.filename, tt.args.direction)
				assert.NoError(t, err)

				assert.Equal(t, tt.wantContent, content)
			}
		})
	}
}

func FuzzLocalStorage_Write(f *testing.F) {
	f.Fuzz(func(t *testing.T, value string) {
		s := LocalStorage{}

		fileName := "test"

		err := s.Write(fileName, t.TempDir(), value)
		assert.NoError(t, err)

		got, err := s.Get(fileName, t.TempDir())
		assert.Equal(t, value, got)
	})
}

func TestLocalStorage_GetOpenFile(t *testing.T) {
	type args struct {
		filename  string
		direction string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "write to new file",
			args: args{
				filename:  "test.txt",
				direction: t.TempDir(),
			},
			wantErr: false,
		},
		{
			name: "error on invalid directory",
			args: args{
				filename:  "test.txt",
				direction: "/invalid/" + t.TempDir(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{}
			got, err := s.GetOpenFile(tt.args.filename, tt.args.direction)

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

				fileExist := s.Exists(tt.args.filename, tt.args.direction)
				assert.True(t, fileExist)
			}
		})
	}
}
