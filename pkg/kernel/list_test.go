//go:build unit

package kernel

import (
	"reflect"
	"testing"
	"time"

	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/testutil"
)

func TestList(t *testing.T) {
	tests := []struct {
		name               string
		want               *connect.Connections
		wantErr            bool
		isCreateConnection bool
	}{
		{
			name:               "success - get empty connections",
			want:               &connect.Connections{},
			wantErr:            false,
			isCreateConnection: false,
		},
		{
			name:               "success - create empty connections",
			want:               &connect.Connections{},
			wantErr:            false,
			isCreateConnection: true,
		},
		{
			name: "success - get exists connections",
			want: &connect.Connections{
				Connects: []connect.Connect{
					{
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
					{
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
					{
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
