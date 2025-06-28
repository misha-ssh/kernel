package store

import (
	"reflect"
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/setup"
	"github.com/ssh-connection-manager/kernel/v2/testutil"
)

func TestGetConnections(t *testing.T) {
	tests := []struct {
		name    string
		want    *connect.Connections
		wantErr bool
	}{
		{
			name:    "get empty connections",
			want:    &connect.Connections{},
			wantErr: false,
		},
	}

	if err := testutil.RemoveFileConnections(); err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
