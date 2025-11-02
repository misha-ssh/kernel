//go:build unit

package crypto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecrypt(t *testing.T) {
	type args struct {
		plaintext string
		key       string
	}
	tests := []struct {
		name        string
		args        args
		want        string
		wantErr     bool
		generateKey bool
	}{
		{
			name: "success",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:        "hello world",
			wantErr:     false,
			generateKey: true,
		},
		{
			name: "empty (string) plaintext",
			args: args{
				plaintext: "",
				key:       "",
			},
			want:        "",
			wantErr:     false,
			generateKey: true,
		},
		{
			name: "empty key",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:        "hello world",
			wantErr:     true,
			generateKey: false,
		},
		{
			name: "long size key",
			args: args{
				plaintext: "hello world",
				key:       "32-byte-long-key-1234567890123456",
			},
			want:        "hello world",
			wantErr:     true,
			generateKey: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			if tt.generateKey {
				tt.args.key, err = GenerateKey()
				require.NoError(t, err)
			}

			cryptText, err := Encrypt(tt.args.plaintext, tt.args.key)
			require.Equal(t, tt.wantErr, err != nil)

			got, err := Decrypt(cryptText, tt.args.key)
			require.Equal(t, tt.wantErr, err != nil)

			if !tt.wantErr {
				require.Equal(t, got, tt.want)
			}
		})
	}
}

func TestEncrypt(t *testing.T) {
	type args struct {
		plaintext string
		key       string
	}
	tests := []struct {
		name        string
		args        args
		want        string
		wantErr     bool
		generateKey bool
	}{
		{
			name: "success",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:        "hello world",
			wantErr:     false,
			generateKey: true,
		},
		{
			name: "empty (string) plaintext",
			args: args{
				plaintext: "",
				key:       "",
			},
			want:        "",
			wantErr:     false,
			generateKey: true,
		},
		{
			name: "empty key",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:        "hello world",
			wantErr:     true,
			generateKey: false,
		},
		{
			name: "long size key",
			args: args{
				plaintext: "hello world",
				key:       "32-byte-long-key-1234567890123456",
			},
			want:        "hello world",
			wantErr:     true,
			generateKey: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			if tt.generateKey {
				tt.args.key, err = GenerateKey()
				require.NoError(t, err)
			}

			encryptText, err := Encrypt(tt.args.plaintext, tt.args.key)
			require.Equal(t, tt.wantErr, err != nil)

			got, err := Decrypt(encryptText, tt.args.key)
			require.Equal(t, tt.wantErr, err != nil)

			if !tt.wantErr {
				require.Equal(t, got, tt.want)
			}
		})
	}
}

func TestGenerateKey(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "default test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GenerateKey()
			require.Equal(t, tt.wantErr, err != nil)
			require.True(t, len(key) == SizeKey)
		})
	}
}
