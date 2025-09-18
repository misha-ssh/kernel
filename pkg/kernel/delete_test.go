package kernel

import (
	"testing"

	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/testutil"
)

func TestDelete(t *testing.T) {
	tempDir := t.TempDir()

	pathToPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		connection *connect.Connect
		isCreate   bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - delete connection",
			args: args{
				connection: &connect.Connect{
					Alias:      tempDir,
					Login:      "test",
					Address:    "test",
					Password:   "test",
					Type:       connect.TypeSSH,
					CreatedAt:  "time",
					UpdatedAt:  "time",
					SshOptions: &connect.SshOptions{},
				},
				isCreate: true,
			},
			wantErr: false,
		},
		{
			name: "fail - not found connection",
			args: args{
				connection: &connect.Connect{
					Alias:      "notFoundAlias",
					Login:      "test",
					Address:    "test",
					Password:   "test",
					Type:       connect.TypeSSH,
					CreatedAt:  "time",
					UpdatedAt:  "time",
					SshOptions: &connect.SshOptions{},
				},
				isCreate: false,
			},
			wantErr: true,
		},
		{
			name: "success - delete connection with private key",
			args: args{
				connection: &connect.Connect{
					Alias:     tempDir,
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
				isCreate: true,
			},
			wantErr: false,
		},
		{
			name: "success - delete connection with empty private key",
			args: args{
				connection: &connect.Connect{
					Alias:     tempDir,
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
				isCreate: true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.isCreate {
				if err = Create(tt.args.connection); err != nil {
					t.Errorf("Create connection error = %v", err)
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
