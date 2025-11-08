//go:build unit

package store

import (
	"testing"

	"github.com/misha-ssh/kernel/internal/setup"
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/testutil"
	"github.com/stretchr/testify/require"
)

func TestStore_ValidatePrivateKey(t *testing.T) {
	tempDir := t.TempDir()

	passphrase := "passphrase"

	pathToPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	require.NoError(t, err)

	pathToPrivateKeyWithPass, err := testutil.CreatePrivateKeyWithPass(tempDir, passphrase)
	require.NoError(t, err)

	pathToInvalidKey, err := testutil.CreateInvalidPrivateKey(tempDir)
	require.NoError(t, err)

	type args struct {
		privateKey string
		passphrase string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - validate data private key",
			args: args{
				privateKey: pathToPrivateKey,
				passphrase: "",
			},
			wantErr: false,
		},
		{
			name: "success - validate data private key with passphrase",
			args: args{
				privateKey: pathToPrivateKeyWithPass,
				passphrase: passphrase,
			},
			wantErr: false,
		},
		{
			name: "fail - invalid data private key",
			args: args{
				privateKey: pathToInvalidKey,
				passphrase: "",
			},
			wantErr: true,
		},
		{
			name: "fail - invalid passphrase data",
			args: args{
				privateKey: pathToPrivateKeyWithPass,
				passphrase: "invalid passphrase",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			direction, filename := storage.GetDirectionAndFilename(tt.args.privateKey)
			privateKey, err := storage.Get(direction, filename)
			require.NoError(t, err)

			err = validatePrivateKey(privateKey, tt.args.passphrase)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestDeletePrivateKey(t *testing.T) {
	tempDir := t.TempDir()

	passphrase := "passphrase"

	pathToPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	require.NoError(t, err)

	pathToPrivateKeyWithPass, err := testutil.CreatePrivateKeyWithPass(tempDir, passphrase)
	require.NoError(t, err)

	pathToInvalidKey, err := testutil.CreateInvalidPrivateKey(tempDir)
	require.NoError(t, err)

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
			name: "success - delete with key and passphrase",
			args: args{
				connection: &connect.Connect{
					Alias: "test_alias",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToPrivateKeyWithPass,
						Passphrase: passphrase,
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
			require.Equal(t, tt.wantErr, err != nil)

			tt.args.connection.SshOptions.PrivateKey = pathSavedPrivateKey

			err = DeletePrivateKey(tt.args.connection)
			require.Equal(t, tt.wantErr, err != nil)

			require.False(t, storage.Exists(storage.GetPrivateKeysDir(), tt.args.connection.Alias))
		})
	}
}

func TestSavePrivateKey(t *testing.T) {
	tempDir := t.TempDir()

	passphrase := "passphrase"

	pathToPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	require.NoError(t, err)

	pathToPrivateKeyWithPass, err := testutil.CreatePrivateKeyWithPass(tempDir, passphrase)
	require.NoError(t, err)

	pathToInvalidKey, err := testutil.CreateInvalidPrivateKey(tempDir)
	require.NoError(t, err)

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
			name: "success - save private key with passphrase",
			args: args{
				connection: &connect.Connect{
					Alias: "test_alias",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToPrivateKeyWithPass,
						Passphrase: passphrase,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "fail - invalid passphrase",
			args: args{
				connection: &connect.Connect{
					Alias: "test_alias",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToPrivateKeyWithPass,
						Passphrase: "invalid passphrase",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "fail - nonexistent private key",
			args: args{
				connection: &connect.Connect{
					Alias: "test_alias",
					SshOptions: &connect.SshOptions{
						PrivateKey: "non-existent-key",
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
			_, err := SavePrivateKey(tt.args.connection)
			require.Equal(t, tt.wantErr, err != nil)

			if !tt.wantErr {
				require.True(t, storage.Exists(storage.GetPrivateKeysDir(), tt.args.connection.Alias))
			}
		})
	}
}

func TestUpdatePrivateKey(t *testing.T) {
	tempDir := t.TempDir()
	passphrase := "passphrase"

	pathToPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	require.NoError(t, err)

	pathToExtraPrivateKey, err := testutil.CreatePrivateKey(tempDir)
	require.NoError(t, err)

	pathToInvalidKey, err := testutil.CreateInvalidPrivateKey(tempDir)
	require.NoError(t, err)

	pathToPrivateKeyWithPass, err := testutil.CreatePrivateKeyWithPass(tempDir, passphrase)
	require.NoError(t, err)

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
			name: "success - save private key with passphrase",
			args: args{
				connection: &connect.Connect{
					Alias: "test",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToPrivateKeyWithPass,
						Passphrase: passphrase,
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
			name: "fail - invalid passphrase",
			args: args{
				connection: &connect.Connect{
					Alias: "test",
					SshOptions: &connect.SshOptions{
						PrivateKey: pathToPrivateKeyWithPass,
						Passphrase: "invalid passphrase",
					},
				},
			},
			want: func() string {
				return storage.GetFullPath(
					storage.GetPrivateKeysDir(),
					"test",
				)
			},
			wantErr: true,
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

	setup.Init()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			pathOldPrivateKey := tt.args.connection.SshOptions.PrivateKey

			pathCreatedKey, err := UpdatePrivateKey(tt.args.connection)
			require.Equal(t, tt.wantErr, err != nil)

			if !tt.wantErr {
				require.Equal(t, pathCreatedKey, tt.want())
			}

			if len(pathOldPrivateKey) != 0 && len(pathCreatedKey) != 0 {
				directionOldKey, filenameOldKey := storage.GetDirectionAndFilename(pathOldPrivateKey)
				dataOldKey, err := storage.Get(directionOldKey, filenameOldKey)
				require.NoError(t, err)

				directionCreatedKey, filenameCreatedKey := storage.GetDirectionAndFilename(pathCreatedKey)
				dataCreatedKey, err := storage.Get(directionCreatedKey, filenameCreatedKey)
				require.NoError(t, err)

				require.Equal(t, dataOldKey, dataCreatedKey)
			}

		})
	}
}
