package logger

import (
	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"testing"
)

func TestStorageLogger_Error(t *testing.T) {
	type fields struct {
		Storage storage.Storage
	}
	type args struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := &StorageLogger{
				Storage: tt.fields.Storage,
			}
			sl.Error(tt.args.value)
		})
	}
}

func TestStorageLogger_log(t *testing.T) {
	type fields struct {
		Storage storage.Storage
	}
	type args struct {
		value any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success logging function",
			fields: fields{
				Storage: &storage.LocalStorage{
					BaseDir: "test",
				},
			},
			args: args{
				value: 123,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := &StorageLogger{
				Storage: tt.fields.Storage,
			}
			if err := sl.log(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("log() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
