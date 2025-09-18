package kernel

import (
	"testing"

	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/testutil"
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

	createdConnection := &connect.Connect{
		Alias:      t.TempDir(),
		Login:      "test",
		Address:    "test",
		Password:   "test",
		Type:       connect.TypeSSH,
		CreatedAt:  "time",
		UpdatedAt:  "time",
		SshOptions: &connect.SshOptions{},
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
					Alias:      t.TempDir(),
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
					Alias:      createdConnection.Alias,
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
					Alias:     t.TempDir(),
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
					Alias:     t.TempDir(),
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
					Alias:     t.TempDir(),
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

	if err = Create(createdConnection); err != nil {
		t.Errorf("Create connection error = %v", err)
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
}
