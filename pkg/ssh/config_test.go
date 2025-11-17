package ssh

import (
	"testing"

	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/stretchr/testify/assert"
)

func TestConfig_GetConnections(t *testing.T) {
	tests := []struct {
		name    string
		want    *connect.Connections
		wantErr bool
	}{
		{
			name:    "success - get connections",
			want:    &connect.Connections{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfig()

			got, err := config.GetConnections()
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
