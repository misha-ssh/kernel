package ssh

import (
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/testutil"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/misha-ssh/kernel/pkg/connect"
)

func TestConfig_GetConnections(t *testing.T) {
	tests := []struct {
		name     string
		want     *connect.Connections
		filename string
		wantErr  bool
	}{
		{
			name:     "success - get connections",
			want:     &connect.Connections{},
			filename: "testdata/config",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()

			err := testutil.CreateSSHConfig(tmpDir, tt.filename)
			require.NoError(t, err)

			_, err = testutil.CreatePrivateKey(tmpDir)
			require.NoError(t, err)

			config := &Config{
				LocalStorage: &storage.Local{
					Path: tmpDir,
				},
			}

			got, err := config.GetConnections()
			if !tt.wantErr {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want, got)
		})
	}
}
