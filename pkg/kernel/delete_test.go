package kernel

import (
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/configs/envconst"
	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

func TestDelete(t *testing.T) {
	type args struct {
		connection *connect.Connect
	}
	tests := []struct {
		name                   string
		args                   args
		wantErr                bool
		isCreateConnection     bool
		isDeleteFileConnection bool
	}{
		{
			name: "success delete connection",
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
			},
			wantErr:                false,
			isCreateConnection:     true,
			isDeleteFileConnection: true,
		},
		{
			name: "delete - not found connection",
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
			},
			wantErr:                true,
			isCreateConnection:     false,
			isDeleteFileConnection: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isDeleteFileConnection {
				if err := storage.Delete(storage.GetAppDir(), envconst.FilenameConnections); (err != nil) != tt.wantErr {
					t.Errorf("failed to delete connection %v", err)
				}
			}

			if tt.isCreateConnection {
				if err := Create(tt.args.connection); (err != nil) != tt.wantErr {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			if err := Delete(tt.args.connection); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
