package kernel

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"github.com/ssh-connection-manager/kernel/v2/testutil"
)

func TestDelete(t *testing.T) {
	tempDir := t.TempDir()

	pathToPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		connection *connect.Connect
	}
	tests := []struct {
		name               string
		args               args
		wantErr            bool
		isCreateConnection bool
	}{
		{
			name: "success - delete connection",
			args: args{
				connection: &connect.Connect{
					Alias:      "test",
					Login:      "test",
					Address:    "test",
					Password:   "test",
					Type:       connect.TypeSSH,
					CreatedAt:  "time",
					UpdatedAt:  "time",
					SshOptions: &connect.SshOptions{},
				},
			},
			wantErr:            false,
			isCreateConnection: true,
		},
		{
			name: "fail - not found connection",
			args: args{
				connection: &connect.Connect{
					Alias:      "test",
					Login:      "test",
					Address:    "test",
					Password:   "test",
					Type:       connect.TypeSSH,
					CreatedAt:  "time",
					UpdatedAt:  "time",
					SshOptions: &connect.SshOptions{},
				},
			},
			wantErr:            true,
			isCreateConnection: false,
		},
		{
			name: "success - delete connection with private key",
			args: args{
				connection: &connect.Connect{
					Alias:     "test",
					Login:     "test",
					Address:   "test",
					Password:  "test",
					Type:      connect.TypeSSH,
					CreatedAt: "time",
					UpdatedAt: "time",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToPrivateKey,
					},
				},
			},
			wantErr:            false,
			isCreateConnection: true,
		},
		{
			name: "success - delete connection with empty private key",
			args: args{
				connection: &connect.Connect{
					Alias:     "test",
					Login:     "test",
					Address:   "test",
					Password:  "test",
					Type:      connect.TypeSSH,
					CreatedAt: "time",
					UpdatedAt: "time",
					SshOptions: &connect.SshOptions{
						PrivateKey: "",
					},
				},
			},
			wantErr:            false,
			isCreateConnection: true,
		},
	}

	if err = testutil.RemoveFileConnections(); err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isCreateConnection {
				if err := Create(tt.args.connection); err != nil {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			if err := Delete(tt.args.connection); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(tt.args.connection.SshOptions.PrivateKey) != 0 {
				direction, filename := storage.GetDirectionAndFilename(tt.args.connection.SshOptions.PrivateKey)
				if storage.Exists(direction, filename) {
					t.Errorf("failed to check connection %v", err)
				}
			}
		})
	}

	if err = testutil.RemoveDirectionPrivateKey(); err != nil {
		t.Fatal(err)
	}
}
