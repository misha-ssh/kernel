//go:build unit

package store

import (
	"reflect"
	"testing"

	"github.com/misha-ssh/kernel/internal/setup"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/stretchr/testify/require"
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup.Init()

			got, err := GetConnections()
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, reflect.TypeOf(got), reflect.TypeOf(tt.want))
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

			err := SetConnections(tt.args.connections)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}
