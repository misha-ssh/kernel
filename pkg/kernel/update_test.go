package kernel

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/testutil"
)

func TestUpdate(t *testing.T) {
	tempDir := t.TempDir()

	pathToPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		connection *connect.Connect
		oldAlias   string
	}
	tests := []struct {
		name               string
		args               args
		wantErr            bool
		isCreateConnection bool
	}{
		{
			name: "success update - update on default old value",
			args: args{
				connection: &connect.Connect{
					Alias:      "test",
					Login:      "test",
					Address:    "test",
					Password:   "test",
					Type:       connect.TypeSSH,
					CreatedAt:  "time",
					UpdatedAt:  "time",
					SshOptions: nil,
				},
				oldAlias: "test",
			},
			wantErr:            false,
			isCreateConnection: true,
		},
		{
			name: "not found connection - get exist connect and get not exists old alias",
			args: args{
				connection: &connect.Connect{
					Alias:      "test",
					Login:      "test",
					Address:    "test",
					Password:   "test",
					Type:       connect.TypeSSH,
					CreatedAt:  "time",
					UpdatedAt:  "time",
					SshOptions: nil,
				},
				oldAlias: "notFoundAlias",
			},
			wantErr:            true,
			isCreateConnection: false,
		},
		{
			name: "update values by exist connection",
			args: args{
				connection: &connect.Connect{
					Alias:      "update",
					Login:      "update",
					Address:    "update",
					Password:   "update",
					Type:       connect.TypeSSH,
					CreatedAt:  "update",
					UpdatedAt:  "update",
					SshOptions: nil,
				},
				oldAlias: "test",
			},
			wantErr:            false,
			isCreateConnection: false,
		},
		{
			name: "a",
			args: args{
				connection: &connect.Connect{
					Alias:      "update",
					Login:      "update",
					Address:    "update",
					Password:   "update",
					Type:       connect.TypeSSH,
					CreatedAt:  "update",
					UpdatedAt:  "update",
					SshOptions: &connect.SshOptions{PrivateKey: pathToPrivateKey},
				},
				oldAlias: "test",
			},
			wantErr:            false,
			isCreateConnection: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isCreateConnection {
				if err := Create(tt.args.connection); (err != nil) != tt.wantErr {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			if err := Update(tt.args.connection, tt.args.oldAlias); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
