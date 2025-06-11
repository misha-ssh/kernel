package store

import (
	"github.com/ssh-connection-manager/kernel/v2/configs/envconst"
	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/setup"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"reflect"
	"testing"
)

func TestGetConnections(t *testing.T) {
	tests := []struct {
		name                   string
		want                   *connect.Connections
		wantErr                bool
		isDeleteFileConnection bool
	}{
		{
			name:                   "get empty connections",
			want:                   &connect.Connections{},
			wantErr:                false,
			isDeleteFileConnection: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isDeleteFileConnection {
				err := storage.Delete(storage.GetAppDir(), envconst.FilenameConnections)
				if err != nil {
					t.Errorf("failed to delete connection %v", err)
				}
			}

			setup.Init()

			got, err := GetConnections()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConnections() error = %v, wantErr %v", err, tt.wantErr)
			}

			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("GetConnections() got: %v != want: %v", got, tt.want)
			}
		})
	}
}

func TestSetConnections(t *testing.T) {
	type args struct {
		connections *connect.Connections
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success set connections",
			args: args{
				connections: &connect.Connections{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup.Init()

			if err := SetConnections(tt.args.connections); (err != nil) != tt.wantErr {
				t.Errorf("SetConnections() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
