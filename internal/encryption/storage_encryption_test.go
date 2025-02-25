package encryption

import (
	"crypto/cipher"
	"github.com/ssh-connection-manager/kernel/v2/pkg/storage"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestStorageEncryption_Decrypt(t *testing.T) {
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

			s := &StorageEncryption{}

			if tt.keyGenerate {
				tt.args.key, err = s.GenerateKey()
			}

			cryptText, err := s.Encrypt(tt.args.plaintext, tt.args.key)
			got, err := s.Decrypt(cryptText, tt.args.key)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestStorageEncryption_Encrypt(t *testing.T) {
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

			s := &StorageEncryption{}

			if tt.keyGenerate {
				tt.args.key, err = s.GenerateKey()
			}

			encryptText, err := s.Encrypt(tt.args.plaintext, tt.args.key)
			got, err := s.Decrypt(encryptText, tt.args.key)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestStorageEncryption_GenerateKey(t *testing.T) {
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
			s := &StorageEncryption{}
			got, err := s.GenerateKey()

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

func TestStorageEncryption_GetKey(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(*storage.MockStorage)
		want      string
		wantErr   bool
	}{
		{
			name: "success",
			setupMock: func(m *storage.MockStorage) {
				m.On("Exists", FileName).Return(true, nil)
				m.On("Get", FileName).Return("key", nil)
			},
			want:    "key",
			wantErr: false,
		},
		{
			name: "success not exists file",
			setupMock: func(m *storage.MockStorage) {
				m.On("Exists", FileName).Return(false, nil)
				m.On("Write", FileName, "2").Return(nil)
			},
			want:    "key",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := new(storage.MockStorage)
			tt.setupMock(mockStorage)

			s := &StorageEncryption{}

			got, err := s.GetKey(mockStorage)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			mockStorage.AssertExpectations(t)
		})
	}
}

func TestStorageEncryption_getGcm(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    cipher.AEAD
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StorageEncryption{}
			got, err := s.getGcm(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("getGcm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGcm() got = %v, want %v", got, tt.want)
			}
		})
	}
}
