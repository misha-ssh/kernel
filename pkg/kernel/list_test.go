package kernel

import (
	"reflect"
	"testing"

	"github.com/misha-ssh/kernel/pkg/connect"
)

func TestList(t *testing.T) {
	tests := []struct {
		name               string
		want               *connect.Connections
		wantErr            bool
		isCreateConnection bool
	}{
		{
			name:               "list connections - get empty connections",
			want:               &connect.Connections{},
			wantErr:            false,
			isCreateConnection: false,
		},
		{
			name:               "list connections - create empty connections",
			want:               &connect.Connections{},
			wantErr:            false,
			isCreateConnection: true,
		},
		{
			name: "list connections - create empty connections",
			want: &connect.Connections{
				Connects: []connect.Connect{
					{
						Alias:      t.TempDir(),
						Login:      "test",
						Address:    "test",
						Password:   "test",
						Type:       connect.TypeSSH,
						CreatedAt:  "time",
						UpdatedAt:  "time",
						SshOptions: &connect.SshOptions{},
					},
					{
						Alias:      t.TempDir(),
						Login:      "test",
						Address:    "test",
						Password:   "test",
						Type:       connect.TypeSSH,
						CreatedAt:  "time",
						UpdatedAt:  "time",
						SshOptions: &connect.SshOptions{},
					},
					{
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
			},
			wantErr:            false,
			isCreateConnection: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
