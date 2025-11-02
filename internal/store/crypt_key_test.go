//go:build unit

package store

import (
	"os/user"
	"testing"

	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
)

func TestGetCryptKey(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success get crypt key",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCryptKey()
			require.Equal(t, tt.wantErr, err != nil)

			currentUser, _ := user.Current()
			want, _ := keyring.Get(envconst.NameServiceCryptKey, currentUser.Username)

			require.Equal(t, got, want)
		})
	}
}
