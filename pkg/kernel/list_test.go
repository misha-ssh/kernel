package kernel

import (
	"reflect"
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/configs/envconst"
	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

func TestList(t *testing.T) {
	tests := []struct {
		name                   string
		want                   *connect.Connections
		wantErr                bool
		isDeleteFileConnection bool
		isCreateConnection     bool
	}{
		{
			name:                   "list connections - get empty connections",
			want:                   &connect.Connections{},
			wantErr:                false,
			isDeleteFileConnection: true,
			isCreateConnection:     false,
		},
		{
			name:                   "list connections - create empty connections",
			want:                   &connect.Connections{},
			wantErr:                false,
			isDeleteFileConnection: true,
			isCreateConnection:     true,
		},
		{
			name: "list connections - create empty connections",
			want: &connect.Connections{
				Connects: []connect.Connect{
					connect.Connect{
						Alias:      "test",
						Login:      "test",
						Address:    "test",
						Password:   "test",
						Type:       connect.TypeSSH,
						CreatedAt:  "time",
						UpdatedAt:  "time",
						SshOptions: nil,
					},
					connect.Connect{
						Alias:      "test 2",
						Login:      "test",
						Address:    "test",
						Password:   "test",
						Type:       connect.TypeSSH,
						CreatedAt:  "time",
						UpdatedAt:  "time",
						SshOptions: nil,
					},
					connect.Connect{
						Alias:      "test 3",
						Login:      "test",
						Address:    "test",
						Password:   "test",
						Type:       connect.TypeSSH,
						CreatedAt:  "time",
						UpdatedAt:  "time",
						SshOptions: nil,
					},
				},
			},
			wantErr:                false,
			isDeleteFileConnection: true,
			isCreateConnection:     true,
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
				for _, connection := range tt.want.Connects {
					if err := Create(&connection); (err != nil) != tt.wantErr {
						t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
					}
				}
			}

			got, err := List()
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
			}

			if reflect.TypeOf(got).String() != reflect.TypeOf(tt.want).String() {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}
