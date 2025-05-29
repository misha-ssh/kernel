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
		name     string
		args     args
		want     string
		wantErr  bool
		isGetKey bool
	}{
		{
			name: "success",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:     "hello world",
			wantErr:  false,
			isGetKey: true,
		},
		{
			name: "empty (string) plaintext",
			args: args{
				plaintext: "",
				key:       "",
			},
			want:     "",
			wantErr:  false,
			isGetKey: true,
		},
		{
			name: "empty key",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:     "hello world",
			wantErr:  true,
			isGetKey: false,
		},
		{
			name: "long size key",
			args: args{
				plaintext: "hello world",
				key:       "32-byte-long-key-1234567890123456",
			},
			want:     "hello world",
			wantErr:  true,
			isGetKey: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			if tt.isGetKey {
				password := "password"
				tt.args.key, err = GetKey(password)
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
		name     string
		args     args
		want     string
		wantErr  bool
		isGetKey bool
	}{
		{
			name: "success",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:     "hello world",
			wantErr:  false,
			isGetKey: true,
		},
		{
			name: "empty (string) plaintext",
			args: args{
				plaintext: "",
				key:       "",
			},
			want:     "",
			wantErr:  false,
			isGetKey: true,
		},
		{
			name: "empty key",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:     "hello world",
			wantErr:  true,
			isGetKey: false,
		},
		{
			name: "long size key",
			args: args{
				plaintext: "hello world",
				key:       "32-byte-long-key-1234567890123456",
			},
			want:     "hello world",
			wantErr:  true,
			isGetKey: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			if tt.isGetKey {
				password := "password"
				tt.args.key, err = GetKey(password)
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

func TestGetKey(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "default test",
			args: args{
				password: "password",
			},
			wantErr: false,
		},
		{
			name: "empty password",
			args: args{
				password: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetKey(tt.args.password)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
