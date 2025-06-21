package store

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/crypto"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

func createInvalidPrivateKey(direction string) (string, error) {
	filenameInvalidKey := "invalid"
	err := storage.Write(direction, filenameInvalidKey, "")
	if err != nil {
		return "", err
	}

	return filepath.Join(direction, filenameInvalidKey), nil
}

func createPrivateKey(direction string) (string, error) {
	filenameKey := "key"

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", err
	}

	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	privatePEM := pem.EncodeToMemory(&privBlock)

	err = storage.Write(direction, filenameKey, string(privatePEM))
	if err != nil {
		return "", err
	}

	return filepath.Join(direction, filenameKey), nil
}

func TestDeletePrivateKey(t *testing.T) {
	type args struct {
		connection *connect.Connect
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeletePrivateKey(tt.args.connection); (err != nil) != tt.wantErr {
				t.Errorf("DeletePrivateKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetPrivateKey(t *testing.T) {
	type args struct {
		connection *connect.Connect
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetPrivateKey(tt.args.connection); (err != nil) != tt.wantErr {
				t.Errorf("GetPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSavePrivateKey(t *testing.T) {
	tempDir := t.TempDir()

	pathToPrivateKey, err := createPrivateKey(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	pathToInvalidKey, err := createInvalidPrivateKey(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		connection *connect.Connect
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "successful save",
			args: args{
				connection: &connect.Connect{
					Alias: t.TempDir(),
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToPrivateKey,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "nonexistent private key",
			args: args{
				connection: &connect.Connect{
					Alias: t.TempDir(),
					SshOptions: &connect.SshOptions{
						PrivateKey: filepath.Join(tempDir, "non-existent-key"),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid private key",
			args: args{
				connection: &connect.Connect{
					Alias: t.TempDir(),
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToInvalidKey,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "empty alias",
			args: args{
				connection: &connect.Connect{
					Alias: "",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToPrivateKey,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "empty private key",
			args: args{
				connection: &connect.Connect{
					Alias: t.TempDir(),
					SshOptions: &connect.SshOptions{
						PrivateKey: "",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			savedPathPrivateKey, err := SavePrivateKey(tt.args.connection)

			if (err != nil) != tt.wantErr {
				t.Errorf("SavePrivateKey() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if !storage.Exists(storage.GetPrivateKeysDir(), tt.args.connection.Alias) {
					t.Errorf("SavePrivateKey() dont create file error = %v, wantErr %v", err, tt.wantErr)
				}

				cryptKey, err := GetCryptKey()
				if err != nil {
					t.Errorf("GetCryptKey() error = %v", err)
				}

				directionSavedPathPrivateKey := filepath.Dir(savedPathPrivateKey)
				filenameSavedPathPrivateKey := filepath.Base(savedPathPrivateKey)

				dataSavedPrivateKey, err := storage.Get(directionSavedPathPrivateKey, filenameSavedPathPrivateKey)
				if err != nil {
					t.Errorf("Get() error = %v", err)
				}

				decryptDataSavedPrivateKey, err := crypto.Decrypt(dataSavedPrivateKey, cryptKey)
				if err != nil {
					t.Errorf("Decrypt() error = %v", err)
				}

				directionPathPrivateKey := filepath.Dir(tt.args.connection.SshOptions.PrivateKey)
				filenamePathPrivateKey := filepath.Base(tt.args.connection.SshOptions.PrivateKey)

				dataPrivateKey, err := storage.Get(directionPathPrivateKey, filenamePathPrivateKey)
				if err != nil {
					t.Errorf("Get() error = %v", err)
				}

				if !reflect.DeepEqual(decryptDataSavedPrivateKey, dataPrivateKey) {
					t.Error("saved private key != saved private key")
				}
			}
		})
	}
}

func Test_validatePrivateKey(t *testing.T) {
	type args struct {
		privateKey string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePrivateKey(tt.args.privateKey); (err != nil) != tt.wantErr {
				t.Errorf("validatePrivateKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
