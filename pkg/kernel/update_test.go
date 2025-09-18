package kernel

import (
	"testing"

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
		Alias:      tempDir,
		Login:      "test",
		Address:    "test",
		Password:   "test",
		Type:       connect.TypeSSH,
		CreatedAt:  "time",
		UpdatedAt:  "time",
		SshOptions: &connect.SshOptions{},
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
					Alias:      createdConnection.Alias,
					Login:      "test",
					Address:    "test",
					Password:   "test",
					Type:       connect.TypeSSH,
					CreatedAt:  "time",
					UpdatedAt:  "time",
					SshOptions: &connect.SshOptions{},
				},
				oldAlias: createdConnection.Alias,
			},
			wantErr: false,
		},
		{
			name: "fail - get exist connect and get not exists old alias",
			args: args{
				connection: &connect.Connect{
					Alias:      createdConnection.Alias,
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
			wantErr: true,
		},
		{
			name: "success - update values by exist connection",
			args: args{
				connection: &connect.Connect{
					Alias:      createdConnection.Alias,
					Login:      "test2",
					Address:    "test2",
					Password:   "test2",
					Type:       connect.TypeSSH,
					CreatedAt:  "test2",
					UpdatedAt:  "test2",
					SshOptions: &connect.SshOptions{},
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
					CreatedAt: "test2",
					UpdatedAt: "test2",
					SshOptions: &connect.SshOptions{
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
					CreatedAt: "test2",
					UpdatedAt: "test2",
					SshOptions: &connect.SshOptions{
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
					CreatedAt: "test2",
					UpdatedAt: "test2",
					SshOptions: &connect.SshOptions{
						PrivateKey: "",
					},
				},
				oldAlias: createdConnection.Alias,
			},
			wantErr: false,
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

	if err = testutil.RemoveDirectionPrivateKey(); err != nil {
		t.Fatal(err)
	}
}
