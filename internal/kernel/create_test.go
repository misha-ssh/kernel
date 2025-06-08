package kernel

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
)

func TestCreate(t *testing.T) {
	type args struct {
		connect *connect.Connect
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create success data",
			args: args{
				connect: &connect.Connect{
					Alias:      t.TempDir(),
					Login:      "test",
					Address:    "test",
					Password:   "test",
					Type:       connect.TypeSSH,
					CreatedAt:  "time",
					UpdatedAt:  "time",
					SshOptions: nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create(tt.args.connect); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
