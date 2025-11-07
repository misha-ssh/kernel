//go:build unit

package kernel

import (
	"testing"
	"time"

	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/testutil"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	tempDir := t.TempDir()

	pathToPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	require.NoError(t, err)

	pathToInvalidKey, err := testutil.CreateInvalidPrivateKey(tempDir)
	require.NoError(t, err)

	type args struct {
		connect *connect.Connect
	}

	createdConnection := &connect.Connect{
		Alias:     testutil.RandomString(),
		Login:     "test",
		Address:   "test",
		Password:  "test",
		Type:      connect.TypeSSH,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
		SshOptions: &connect.SshOptions{
			Port: 22,
		},
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
					Alias:     testutil.RandomString(),
					Login:     "test",
					Address:   "test",
					Password:  "test",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port: 22,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "fail - exist alias",
			args: args{
				connect: &connect.Connect{
					Alias:     createdConnection.Alias,
					Login:     "test",
					Address:   "test",
					Password:  "test",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port: 22,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "success - add ssh options",
			args: args{
				connect: &connect.Connect{
					Alias:     testutil.RandomString(),
					Login:     "test",
					Address:   "test",
					Password:  "test",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port:       22,
						PrivateKey: "",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success - save private key with password",
			args: args{
				connect: &connect.Connect{
					Alias:     testutil.RandomString(),
					Login:     "test",
					Address:   "test",
					Password:  "test",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port:       22,
						PrivateKey: pathToPrivateKey,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success - save private key",
			args: args{
				connect: &connect.Connect{
					Alias:     testutil.RandomString(),
					Login:     "test",
					Address:   "test",
					Password:  "",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port:       22,
						PrivateKey: pathToPrivateKey,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success - save private key with pass",
			args: args{
				connect: &connect.Connect{
					Alias:     testutil.RandomString(),
					Login:     "test",
					Address:   "test",
					Password:  "",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port:       22,
						PrivateKey: pathToPrivateKey,
						Passphrase: "password",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "fail - dont valid private key",
			args: args{
				connect: &connect.Connect{
					Alias:     testutil.RandomString(),
					Login:     "test",
					Address:   "test",
					Password:  "test",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port:       22,
						PrivateKey: pathToInvalidKey,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fail - empty alias",
			args: args{
				connect: &connect.Connect{
					Alias:     "",
					Login:     "test",
					Address:   "test",
					Password:  "test",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port: 22,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fail - alias is invalid with /",
			args: args{
				connect: &connect.Connect{
					Alias:     "test/alias",
					Login:     "test",
					Address:   "test",
					Password:  "test",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port: 22,
					},
				},
			},
			wantErr: true,
		},
	}

	require.NoError(t, Create(createdConnection))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				require.Error(t, Create(tt.args.connect))
			} else {
				require.NoError(t, Create(tt.args.connect))
			}

			if len(tt.args.connect.SshOptions.PrivateKey) != 0 {
				direction, filename := storage.GetDirectionAndFilename(tt.args.connect.SshOptions.PrivateKey)
				require.True(t, storage.Exists(direction, filename))
			}
		})
	}
}
