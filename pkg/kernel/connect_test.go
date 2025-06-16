package kernel

import (
	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"testing"
)

func TestConnect(t *testing.T) {
	type args struct {
		connection *connect.Connect
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Connect(tt.args.connection); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
