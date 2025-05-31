package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_Decrypt(t *testing.T) {
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
			}

			cryptText, err := Encrypt(tt.args.plaintext, tt.args.key)
			got, err := Decrypt(cryptText, tt.args.key)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestStorage_Encrypt(t *testing.T) {
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
			}

			encryptText, err := Encrypt(tt.args.plaintext, tt.args.key)
			got, err := Decrypt(encryptText, tt.args.key)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
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

			if len(key) < SizeKey {
				t.Errorf("key is invalid error = %v, wantErr %v", err, tt.wantErr)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
