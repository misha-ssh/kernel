package store

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"reflect"
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/connect"
	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
)

func createInvalidPrivateKey(direction string) (string, error) {
	filenameInvalidKey := "invalid"
	err := storage.Write(direction, filenameInvalidKey, "")
	if err != nil {
		return "", err
	}

	return storage.GetFullPath(direction, filenameInvalidKey), nil
}

func createPrivateKey(direction string) (string, error) {
	filenameKey := "key"

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", err
	}

	privateDER := x509.MarshalPKCS1PrivateKey(privateKey)

	privateBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateDER,
	}

	privatePEM := pem.EncodeToMemory(&privateBlock)

	err = storage.Write(direction, filenameKey, string(privatePEM))
	if err != nil {
		return "", err
	}

	return storage.GetFullPath(direction, filenameKey), nil
}

func TestDeletePrivateKey(t *testing.T) {
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
			name: "successful delete",
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
			name: "invalid private key - delete key",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pathSavedPrivateKey, err := SavePrivateKey(tt.args.connection)
			if (err != nil) != tt.wantErr {
				t.Errorf("SavePrivateKey() error = %v, wantErr %v", err, tt.wantErr)
			}

			tt.args.connection.SshOptions.PrivateKey = pathSavedPrivateKey

			if err := DeletePrivateKey(tt.args.connection); (err != nil) != tt.wantErr {
				t.Errorf("DeletePrivateKey() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && storage.Exists(storage.GetPrivateKeysDir(), tt.args.connection.Alias) {
				t.Errorf("SavePrivateKey() dont create file error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	err = os.RemoveAll(storage.GetPrivateKeysDir())
	if err != nil {
		t.Fatal(err)
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
						PrivateKey: storage.GetFullPath(tempDir, "non-existent-key"),
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

				directionSavedPrivateKey, filenameSavedPrivateKey := storage.GetDirectionAndFilename(savedPathPrivateKey)
				dataSavedPrivateKey, err := storage.Get(directionSavedPrivateKey, filenameSavedPrivateKey)
				if err != nil {
					t.Errorf("Get() error = %v", err)
				}

				directionPrivateKey, filenamePrivateKey := storage.GetDirectionAndFilename(tt.args.connection.SshOptions.PrivateKey)
				dataPrivateKey, err := storage.Get(directionPrivateKey, filenamePrivateKey)
				if err != nil {
					t.Errorf("Get() error = %v", err)
				}

				if !reflect.DeepEqual(dataSavedPrivateKey, dataPrivateKey) {
					t.Error("saved private key != private key")
				}
			}
		})
	}

	err = os.RemoveAll(storage.GetPrivateKeysDir())
	if err != nil {
		t.Fatal(err)
	}
}

func Test_validatePrivateKey(t *testing.T) {
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
		privateKey string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "successful validate data private key",
			args: args{
				privateKey: pathToPrivateKey,
			},
			wantErr: false,
		},
		{
			name: "invalid data private key",
			args: args{
				privateKey: pathToInvalidKey,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			direction, filename := storage.GetDirectionAndFilename(tt.args.privateKey)
			privateKey, err := storage.Get(direction, filename)
			if err != nil {
				t.Errorf("Get() error = %v", err)
			}

			if err := validatePrivateKey(privateKey); (err != nil) != tt.wantErr {
				t.Errorf("validatePrivateKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	err = os.RemoveAll(storage.GetPrivateKeysDir())
	if err != nil {
		t.Fatal(err)
	}
}
