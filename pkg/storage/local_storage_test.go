package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalStorage_fullPath(t *testing.T) {
	type fields struct {
		BaseDir string
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "full path with two arg",
			fields: fields{BaseDir: "/test"},
			args:   args{filename: "test.txt"},
			want:   "/test/test.txt",
		},
		{
			name:   "full path with filename",
			fields: fields{BaseDir: ""},
			args:   args{filename: "test.txt"},
			want:   "test.txt",
		},
		{
			name:   "full path with basedir",
			fields: fields{BaseDir: "/test"},
			args:   args{filename: ""},
			want:   "/test",
		},
		{
			name:   "full path with empty args",
			fields: fields{BaseDir: ""},
			args:   args{filename: ""},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{
				BaseDir: tt.fields.BaseDir,
			}

			got := s.fullPath(tt.args.filename)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLocalStorage_Create(t *testing.T) {
	type fields struct {
		BaseDir string
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
			fields:  fields{BaseDir: t.TempDir()},
			args:    args{filename: "test.txt"},
			wantErr: false,
		},
		{
			name:    "create file with empty filename - create dir",
			fields:  fields{BaseDir: t.TempDir()},
			args:    args{filename: ""},
			wantErr: false,
		},
		{
			name:    "create file with empty two args",
			fields:  fields{BaseDir: ""},
			args:    args{filename: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{
				BaseDir: tt.fields.BaseDir,
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
		BaseDir string
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
			fields:       fields{BaseDir: t.TempDir()},
			args:         args{filename: "test.txt"},
			isCreateFile: true,
			wantErr:      false,
		},
		{
			name:         "delete dont created file",
			fields:       fields{BaseDir: t.TempDir()},
			args:         args{filename: "nonexistent.txt"},
			isCreateFile: false,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{
				BaseDir: tt.fields.BaseDir,
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
		BaseDir string
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
			fields:       fields{BaseDir: t.TempDir()},
			args:         args{filename: "test.txt"},
			isCreateFile: true,
			want:         true,
		},
		{
			name:         "is exist dont created file",
			fields:       fields{BaseDir: t.TempDir()},
			args:         args{filename: "nonexistent.txt"},
			isCreateFile: false,
			want:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{
				BaseDir: tt.fields.BaseDir,
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
		BaseDir string
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
			fields:       fields{BaseDir: t.TempDir()},
			args:         args{filename: "test.txt"},
			want:         "test data",
			isCreateFile: true,
			wantErr:      false,
		},
		{
			name:         "get empty file",
			fields:       fields{BaseDir: t.TempDir()},
			args:         args{filename: "nonexistent.txt"},
			want:         "",
			isCreateFile: true,
			wantErr:      false,
		},
		{
			name:         "get dont created file",
			fields:       fields{BaseDir: t.TempDir()},
			args:         args{filename: "nonexistent.txt"},
			want:         "",
			isCreateFile: false,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{
				BaseDir: tt.fields.BaseDir,
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
		BaseDir string
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
			fields:      fields{BaseDir: t.TempDir()},
			args:        args{filename: "test.txt", data: "Hello, World!"},
			wantErr:     false,
			wantContent: "Hello, World!",
		},
		{
			name:        "write to existing file (overwrite)",
			fields:      fields{BaseDir: t.TempDir()},
			args:        args{filename: "test.txt", data: "New content"},
			wantErr:     false,
			wantContent: "New content",
		},
		{
			name:        "write empty data to new file",
			fields:      fields{BaseDir: t.TempDir()},
			args:        args{filename: "empty.txt", data: ""},
			wantErr:     false,
			wantContent: "",
		},
		{
			name:        "write to invalid filename",
			fields:      fields{BaseDir: t.TempDir()},
			args:        args{filename: "", data: "Invalid"},
			wantErr:     true,
			wantContent: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{
				BaseDir: tt.fields.BaseDir,
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

func FuzzLocalStorage_Write(f *testing.F) {
	f.Fuzz(func(t *testing.T, value string) {
		s := LocalStorage{
			BaseDir: t.TempDir(),
		}

		fileName := "test"

		err := s.Write(fileName, value)
		assert.NoError(t, err)

		got, err := s.Get(fileName)
		assert.Equal(t, value, got)
	})
}

func TestLocalStorage_GetOpenFile(t *testing.T) {
	type fields struct {
		BaseDir string
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
			fields:  fields{BaseDir: t.TempDir()},
			args:    args{filename: "test.txt"},
			wantErr: false,
		},
		{
			name:    "error on invalid directory",
			fields:  fields{BaseDir: "/invalid/directory"},
			args:    args{filename: "test.txt"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LocalStorage{
				BaseDir: tt.fields.BaseDir,
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
