package kernel

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/configs/envconst"
	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

func TestCreate(t *testing.T) {
	type args struct {
		connect *connect.Connect
	}
	tests := []struct {
		name                   string
		args                   args
		wantErr                bool
		isDeleteFileConnection bool
	}{
		{
			name: "create connection with random alias",
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
			wantErr:                false,
			isDeleteFileConnection: true,
		},
		{
			name: "create connection with test alias",
			args: args{
				connect: &connect.Connect{
					Alias:      "test",
					Login:      "test",
					Address:    "test",
					Password:   "test",
					Type:       connect.TypeSSH,
					CreatedAt:  "time",
					UpdatedAt:  "time",
					SshOptions: nil,
				},
			},
			wantErr:                false,
			isDeleteFileConnection: false,
		},
		{
			name: "create connection with test alias - get err",
			args: args{
				connect: &connect.Connect{
					Alias:      "test",
					Login:      "test",
					Address:    "test",
					Password:   "test",
					Type:       connect.TypeSSH,
					CreatedAt:  "time",
					UpdatedAt:  "time",
					SshOptions: nil,
				},
			},
			wantErr:                true,
			isDeleteFileConnection: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isDeleteFileConnection {
				err := storage.Delete(storage.GetAppDir(), envconst.FilenameConnections)
				if err != nil {
					t.Errorf("failed to delete connection %v", err)
				}
			}

			if err := Create(tt.args.connect); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
