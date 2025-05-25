package crypto

import (
	"crypto/cipher"
	"github.com/stretchr/testify/assert"
	"testing"
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
		keyGenerate bool
	}{
		{
			name: "success",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:        "hello world",
			wantErr:     false,
			keyGenerate: true,
		},
		{
			name: "empty (string) plaintext",
			args: args{
				plaintext: "",
				key:       "",
			},
			want:        "",
			wantErr:     false,
			keyGenerate: true,
		},
		{
			name: "empty key",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:        "hello world",
			wantErr:     true,
			keyGenerate: false,
		},
		{
			name: "long size key",
			args: args{
				plaintext: "hello world",
				key:       "32-byte-long-key-1234567890123456",
			},
			want:        "hello world",
			wantErr:     true,
			keyGenerate: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			if tt.keyGenerate {
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
		keyGenerate bool
	}{
		{
			name: "success",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:        "hello world",
			wantErr:     false,
			keyGenerate: true,
		},
		{
			name: "empty (string) plaintext",
			args: args{
				plaintext: "",
				key:       "",
			},
			want:        "",
			wantErr:     false,
			keyGenerate: true,
		},
		{
			name: "empty key",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:        "hello world",
			wantErr:     true,
			keyGenerate: false,
		},
		{
			name: "long size key",
			args: args{
				plaintext: "hello world",
				key:       "32-byte-long-key-1234567890123456",
			},
			want:        "hello world",
			wantErr:     true,
			keyGenerate: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			if tt.keyGenerate {
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

func TestStorage_GenerateKey(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateKey()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				assert.NotNil(t, got)
				assert.NotEmpty(t, got)
			}
		})
	}
}

func TestStorage_getGcm(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    cipher.AEAD
		wantErr bool
	}{
		{
			name: "success with 32-byte key",
			args: args{
				key: string(make([]byte, SizeKey)),
			},
			wantErr: false,
		},
		{
			name: "success with 16-byte key",
			args: args{
				key: string(make([]byte, 16)),
			},
			wantErr: false,
		},
		{
			name: "error with invalid key length",
			args: args{
				key: string(make([]byte, 64)),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getGcm(tt.args.key)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}
