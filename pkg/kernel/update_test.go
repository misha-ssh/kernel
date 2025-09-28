//go:build unit

package kernel

import (
	"testing"
	"time"

	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/testutil"
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

	type args struct {
		connection *connect.Connect
		oldAlias   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - update on default old value",
			args: args{
				connection: &connect.Connect{
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
				oldAlias: createdConnection.Alias,
			},
			wantErr: false,
		},
		{
			name: "fail - get exist connect and get not exists old alias",
			args: args{
				connection: &connect.Connect{
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
				oldAlias: "notFoundAlias",
			},
			wantErr: true,
		},
		{
			name: "success - update values by exist connection",
			args: args{
				connection: &connect.Connect{
					Alias:     createdConnection.Alias,
					Login:     "test2",
					Address:   "test2",
					Password:  "test2",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port: 22,
					},
				},
				oldAlias: createdConnection.Alias,
			},
			wantErr: false,
		},
		{
			name: "success - add private key",
			args: args{
				connection: &connect.Connect{
					Alias:     createdConnection.Alias,
					Login:     "test2",
					Address:   "test2",
					Password:  "test2",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port:       22,
						PrivateKey: pathToPrivateKey,
					},
				},
				oldAlias: createdConnection.Alias,
			},
			wantErr: false,
		},
		{
			name: "fail - invalid private key",
			args: args{
				connection: &connect.Connect{
					Alias:     createdConnection.Alias,
					Login:     "test2",
					Address:   "test2",
					Password:  "test2",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port:       22,
						PrivateKey: pathToInvalidKey,
					},
				},
				oldAlias: createdConnection.Alias,
			},
			wantErr: true,
		},
		{
			name: "success - delete private key",
			args: args{
				connection: &connect.Connect{
					Alias:     createdConnection.Alias,
					Login:     "test2",
					Address:   "test2",
					Password:  "test2",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port:       22,
						PrivateKey: "",
					},
				},
				oldAlias: createdConnection.Alias,
			},
			wantErr: false,
		},
		{
			name: "fail - alias is empty",
			args: args{
				connection: &connect.Connect{
					Alias:     "",
					Login:     "test2",
					Address:   "test2",
					Password:  "test2",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port:       22,
						PrivateKey: "",
					},
				},
				oldAlias: createdConnection.Alias,
			},
			wantErr: true,
		},
		{
			name: "fail - alias is invalid with /",
			args: args{
				connection: &connect.Connect{
					Alias:     "/test/",
					Login:     "test2",
					Address:   "test2",
					Password:  "test2",
					Type:      connect.TypeSSH,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
					SshOptions: &connect.SshOptions{
						Port:       22,
						PrivateKey: "",
					},
				},
				oldAlias: createdConnection.Alias,
			},
			wantErr: true,
		},
	}

	if err = Create(createdConnection); err != nil {
		t.Errorf("Create connection error = %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
}
