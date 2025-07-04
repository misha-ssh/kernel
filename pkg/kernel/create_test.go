package kernel

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"github.com/ssh-connection-manager/kernel/v2/pkg/connect"
	"github.com/ssh-connection-manager/kernel/v2/testutil"
)

func TestCreate(t *testing.T) {
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
		connect *connect.Connect
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - create connection",
			args: args{
				connect: &connect.Connect{
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
			wantErr: false,
		},
		{
			name: "fail - exist alias",
			args: args{
				connect: &connect.Connect{
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
			wantErr: true,
		},
		{
			name: "success - add ssh options",
			args: args{
				connect: &connect.Connect{
					Alias:     "test2",
					Login:     "test",
					Address:   "test",
					Password:  "test",
					Type:      connect.TypeSSH,
					CreatedAt: "time",
					UpdatedAt: "time",
					SshOptions: &connect.SshOptions{
						Port:       22,
						PrivateKey: "",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success - save private key",
			args: args{
				connect: &connect.Connect{
					Alias:     "test3",
					Login:     "test",
					Address:   "test",
					Password:  "test",
					Type:      connect.TypeSSH,
					CreatedAt: "time",
					UpdatedAt: "time",
					SshOptions: &connect.SshOptions{
						Port:       22,
						PrivateKey: pathToPrivateKey,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "fail - dont valid private key",
			args: args{
				connect: &connect.Connect{
					Alias:     "test4",
					Login:     "test",
					Address:   "test",
					Password:  "test",
					Type:      connect.TypeSSH,
					CreatedAt: "time",
					UpdatedAt: "time",
					SshOptions: &connect.SshOptions{
						Port:       22,
						PrivateKey: pathToInvalidKey,
					},
				},
			},
			wantErr: true,
		},
	}

	if err = testutil.RemoveFileConnections(); err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create(tt.args.connect); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(tt.args.connect.SshOptions.PrivateKey) != 0 {
				direction, filename := storage.GetDirectionAndFilename(tt.args.connect.SshOptions.PrivateKey)
				if !storage.Exists(direction, filename) {
					t.Errorf("failed to check connection %v", err)
				}
			}
		})
	}

	if err = testutil.RemoveDirectionPrivateKey(); err != nil {
		t.Fatal(err)
	}
}
