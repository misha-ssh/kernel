package store

import (
	"reflect"
	"testing"

	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/testutil"
)

func TestValidatePrivateKey(t *testing.T) {
	tempDir := t.TempDir()

	pathToPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	pathToInvalidKey, err := testutil.CreateInvalidPrivateKey(tempDir)
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
}

func TestDeletePrivateKey(t *testing.T) {
	tempDir := t.TempDir()

	pathToPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	pathToInvalidKey, err := testutil.CreateInvalidPrivateKey(tempDir)
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
					Alias: "test_alias",
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
					Alias: "test_alias",
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
					Alias: "test_alias",
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

			if storage.Exists(storage.GetPrivateKeysDir(), tt.args.connection.Alias) {
				t.Errorf("key exists after delete")
			}
		})
	}
}

func TestSavePrivateKey(t *testing.T) {
	tempDir := t.TempDir()

	pathToPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	pathToInvalidKey, err := testutil.CreateInvalidPrivateKey(tempDir)
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
			name: "success - save private key",
			args: args{
				connection: &connect.Connect{
					Alias: "test_alias",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToPrivateKey,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "fail - nonexistent private key",
			args: args{
				connection: &connect.Connect{
					Alias: "test_alias",
					SshOptions: &connect.SshOptions{
						PrivateKey: storage.GetFullPath(tempDir, "non-existent-key"),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fail - invalid private key",
			args: args{
				connection: &connect.Connect{
					Alias: "test_alias",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToInvalidKey,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fail - empty alias",
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
			name: "fail - empty private key",
			args: args{
				connection: &connect.Connect{
					Alias: "test_alias",
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
}

func TestUpdatePrivateKey(t *testing.T) {
	tempDir := t.TempDir()

	pathToPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	pathToExtraPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	pathToInvalidKey, err := testutil.CreateInvalidPrivateKey(tempDir)
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
		want    func() string
	}{
		{
			name: "successful update - save private key",
			args: args{
				connection: &connect.Connect{
					Alias: "test",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToPrivateKey,
					},
				},
			},
			want: func() string {
				return storage.GetFullPath(
					storage.GetPrivateKeysDir(),
					"test",
				)
			},
			wantErr: false,
		},
		{
			name: "fail update - dont save invalid key",
			args: args{
				connection: &connect.Connect{
					Alias: "test",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToInvalidKey,
					},
				},
			},
			want: func() string {
				return pathToInvalidKey
			},
			wantErr: true,
		},
		{
			name: "successful update - get current path private key",
			args: args{
				connection: &connect.Connect{
					Alias: "test",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToPrivateKey,
					},
				},
			},
			want: func() string {
				return storage.GetFullPath(
					storage.GetPrivateKeysDir(),
					"test",
				)
			},
			wantErr: false,
		},
		{
			name: "successful update - get current path but updated private key",
			args: args{
				connection: &connect.Connect{
					Alias: "test",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToExtraPrivateKey,
					},
				},
			},
			want: func() string {
				return storage.GetFullPath(
					storage.GetPrivateKeysDir(),
					"test",
				)
			},
			wantErr: false,
		},
		{
			name: "successful update - delete and set empty private key",
			args: args{
				connection: &connect.Connect{
					Alias: "test",
					SshOptions: &connect.SshOptions{
						PrivateKey: "",
					},
				},
			},
			want: func() string {
				return ""
			},
			wantErr: false,
		},
		{
			name: "successful update - update not exists key, get empty path",
			args: args{
				connection: &connect.Connect{
					Alias: "test2",
					SshOptions: &connect.SshOptions{
						PrivateKey: "",
					},
				},
			},
			want: func() string {
				return ""
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pathOldPrivateKey := tt.args.connection.SshOptions.PrivateKey

			pathCreatedKey, err := UpdatePrivateKey(tt.args.connection)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (pathCreatedKey != tt.want()) != tt.wantErr {
				t.Errorf("UpdatePrivateKey() got = %v, want %v", pathCreatedKey, tt.want())
			}

			if len(pathOldPrivateKey) != 0 && len(pathCreatedKey) != 0 {
				directionOldKey, filenameOldKey := storage.GetDirectionAndFilename(pathOldPrivateKey)
				dataOldKey, err := storage.Get(directionOldKey, filenameOldKey)
				if err != nil {
					t.Error("old key Get() error = ", err)
				}

				directionCreatedKey, filenameCreatedKey := storage.GetDirectionAndFilename(pathCreatedKey)
				dataCreatedKey, err := storage.Get(directionCreatedKey, filenameCreatedKey)
				if err != nil {
					t.Error("old key Get() error = ", err)
				}

				if !reflect.DeepEqual(dataOldKey, dataCreatedKey) {
					t.Errorf("data old key: %v != data created key: %v", pathOldPrivateKey, dataCreatedKey)
				}
			}

		})
	}
}
