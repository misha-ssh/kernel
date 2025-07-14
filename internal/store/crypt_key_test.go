package store

import (
	"os/user"
	"testing"

	"github.com/misha-ssh/kernel/configs/envconst"
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
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCryptKey() error = %v, wantErr %v", err, tt.wantErr)
			}

			currentUser, _ := user.Current()
			want, _ := keyring.Get(envconst.NameServiceCryptKey, currentUser.Username)

			if got != want {
				t.Errorf("GetCryptKey() got = %v, want %v", got, want)
			}
		})
	}
}
