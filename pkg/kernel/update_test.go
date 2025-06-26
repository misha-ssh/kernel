package kernel

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"github.com/ssh-connection-manager/kernel/v2/testutil"
)

func TestUpdate(t *testing.T) {
	tempDir := t.TempDir()

	pathToPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	pathToInvalidKey, err := testutil.CreateInvalidPrivateKey(tempDir)
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
			name: "success - update on default old value",
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
				oldAlias: "test",
			},
			wantErr:            false,
			isCreateConnection: true,
		},
		{
			name: "fail - get exist connect and get not exists old alias",
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
				oldAlias: "notFoundAlias",
			},
			wantErr:            true,
			isCreateConnection: false,
		},
		{
			name: "success - update values by exist connection",
			args: args{
				connection: &connect.Connect{
					Alias:      "test2",
					Login:      "test2",
					Address:    "test2",
					Password:   "test2",
					Type:       connect.TypeSSH,
					CreatedAt:  "test2",
					UpdatedAt:  "test2",
					SshOptions: &connect.SshOptions{},
				},
				oldAlias: "test",
			},
			wantErr:            false,
			isCreateConnection: false,
		},
		{
			name: "success - add private key",
			args: args{
				connection: &connect.Connect{
					Alias:     "test2",
					Login:     "test2",
					Address:   "test2",
					Password:  "test2",
					Type:      connect.TypeSSH,
					CreatedAt: "test2",
					UpdatedAt: "test2",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToPrivateKey,
					},
				},
				oldAlias: "test2",
			},
			wantErr:            false,
			isCreateConnection: false,
		},
		{
			name: "fail - invalid private key",
			args: args{
				connection: &connect.Connect{
					Alias:     "test2",
					Login:     "test2",
					Address:   "test2",
					Password:  "test2",
					Type:      connect.TypeSSH,
					CreatedAt: "test2",
					UpdatedAt: "test2",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToInvalidKey,
					},
				},
				oldAlias: "test2",
			},
			wantErr:            true,
			isCreateConnection: false,
		},
		{
			name: "success - delete private key",
			args: args{
				connection: &connect.Connect{
					Alias:     "test2",
					Login:     "test2",
					Address:   "test2",
					Password:  "test2",
					Type:      connect.TypeSSH,
					CreatedAt: "test2",
					UpdatedAt: "test2",
					SshOptions: &connect.SshOptions{
						PrivateKey: "",
					},
				},
				oldAlias: "test2",
			},
			wantErr:            false,
			isCreateConnection: false,
		},
	}

	if err = testutil.RemoveFileConnections(); err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isCreateConnection {
				if err = Create(tt.args.connection); (err != nil) != tt.wantErr {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			if err = Update(tt.args.connection, tt.args.oldAlias); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}

			direction, filename := storage.GetDirectionAndFilename(tt.args.connection.SshOptions.PrivateKey)

			if !tt.wantErr {
				if len(tt.args.connection.SshOptions.PrivateKey) != 0 {
					if !storage.Exists(direction, filename) {
						t.Error("private key dont exists")
					}
				} else {
					if storage.Exists(direction, filename) {
						t.Error("private key exist but should be removed")
					}
				}
			}
		})
	}

	if err = testutil.RemoveDirectionPrivateKey(); err != nil {
		t.Fatal(err)
	}
}
