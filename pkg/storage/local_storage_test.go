package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalStorage_Create(t *testing.T) {
	type fields struct {
		direction string
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "create file with valida data",
			fields:  fields{direction: t.TempDir()},
			args:    args{filename: "test.txt"},
			wantErr: false,
		},
		{
			name:    "create file with empty filename - create dir",
			fields:  fields{direction: t.TempDir()},
			args:    args{filename: ""},
			wantErr: false,
		},
		{
			name:    "create file with empty two args",
			fields:  fields{direction: ""},
			args:    args{filename: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{
				direction: tt.fields.direction,
			}

			err := s.Create(tt.args.filename)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLocalStorage_Delete(t *testing.T) {
	type fields struct {
		direction string
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		isCreateFile bool
		wantErr      bool
	}{
		{
			name:         "delete existing file",
			fields:       fields{direction: t.TempDir()},
			args:         args{filename: "test.txt"},
			isCreateFile: true,
			wantErr:      false,
		},
		{
			name:         "delete dont created file",
			fields:       fields{direction: t.TempDir()},
			args:         args{filename: "nonexistent.txt"},
			isCreateFile: false,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{
				direction: tt.fields.direction,
			}

			if tt.isCreateFile {
				err := s.Create(tt.args.filename)
				assert.NoError(t, err)
			}

			err := s.Delete(tt.args.filename)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLocalStorage_Exists(t *testing.T) {
	type fields struct {
		direction string
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		isCreateFile bool
		want         bool
	}{
		{
			name:         "is exist created file",
			fields:       fields{direction: t.TempDir()},
			args:         args{filename: "test.txt"},
			isCreateFile: true,
			want:         true,
		},
		{
			name:         "is exist dont created file",
			fields:       fields{direction: t.TempDir()},
			args:         args{filename: "nonexistent.txt"},
			isCreateFile: false,
			want:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{
				direction: tt.fields.direction,
			}

			if tt.isCreateFile {
				err := s.Create(tt.args.filename)
				assert.Nil(t, err)
			}

			got := s.Exists(tt.args.filename)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLocalStorage_Get(t *testing.T) {
	type fields struct {
		direction string
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		want         string
		isCreateFile bool
		wantErr      bool
	}{
		{
			name:         "get dont empty file",
			fields:       fields{direction: t.TempDir()},
			args:         args{filename: "test.txt"},
			want:         "test data",
			isCreateFile: true,
			wantErr:      false,
		},
		{
			name:         "get empty file",
			fields:       fields{direction: t.TempDir()},
			args:         args{filename: "nonexistent.txt"},
			want:         "",
			isCreateFile: true,
			wantErr:      false,
		},
		{
			name:         "get dont created file",
			fields:       fields{direction: t.TempDir()},
			args:         args{filename: "nonexistent.txt"},
			want:         "",
			isCreateFile: false,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{
				direction: tt.fields.direction,
			}

			if tt.isCreateFile {
				err := s.Create(tt.args.filename)
				assert.NoError(t, err)

				err = s.Write(tt.args.filename, tt.want)
				assert.NoError(t, err)
			}

			got, err := s.Get(tt.args.filename)

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
	type fields struct {
		direction string
	}
	type args struct {
		filename string
		data     string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		wantContent string
	}{
		{
			name:        "write to new file",
			fields:      fields{direction: t.TempDir()},
			args:        args{filename: "test.txt", data: "Hello, World!"},
			wantErr:     false,
			wantContent: "Hello, World!",
		},
		{
			name:        "write to existing file (overwrite)",
			fields:      fields{direction: t.TempDir()},
			args:        args{filename: "test.txt", data: "New content"},
			wantErr:     false,
			wantContent: "New content",
		},
		{
			name:        "write empty data to new file",
			fields:      fields{direction: t.TempDir()},
			args:        args{filename: "empty.txt", data: ""},
			wantErr:     false,
			wantContent: "",
		},
		{
			name:        "write to invalid filename",
			fields:      fields{direction: t.TempDir()},
			args:        args{filename: "", data: "Invalid"},
			wantErr:     true,
			wantContent: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{
				direction: tt.fields.direction,
			}

			err := s.Write(tt.args.filename, tt.args.data)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				content, err := s.Get(tt.args.filename)
				assert.NoError(t, err)

				assert.Equal(t, tt.wantContent, content)
			}
		})
	}
}

func TestLocalStorage_GetOpenFile(t *testing.T) {
	type fields struct {
		direction string
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "write to new file",
			fields:  fields{direction: t.TempDir()},
			args:    args{filename: "test.txt"},
			wantErr: false,
		},
		{
			name:    "error on invalid directory",
			fields:  fields{direction: "/invalid/directory"},
			args:    args{filename: "test.txt"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{
				direction: tt.fields.direction,
			}
			got, err := s.GetOpenFile(tt.args.filename)

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

				assert.True(t, s.Exists(tt.args.filename))
			}
		})
	}
}
