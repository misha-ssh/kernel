//go:build unit

package setup

import (
	"os/user"
	"reflect"
	"testing"

	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/configs/envname"
	"github.com/misha-ssh/kernel/internal/config"
	"github.com/misha-ssh/kernel/internal/crypto"
	"github.com/misha-ssh/kernel/internal/logger"
	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "success init",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				require.Nil(t, recover())
			}()

			Init()
		})
	}
}

func TestInitCryptKey(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success set crypt key",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := initCryptKey()
			require.Equal(t, tt.wantErr, err != nil)

			currentUser, err := user.Current()
			require.NoError(t, err)

			username := currentUser.Username
			service := envconst.NameServiceCryptKey
			cryptKey, _ := keyring.Get(service, username)

			require.Equal(t, len(cryptKey), crypto.SizeKey)
		})
	}
}

func TestInitFileConfig(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success created file configs",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.wantErr, initFileConfig() != nil)

			direction := storage.GetAppDir()
			require.True(t, storage.Exists(direction, envconst.FilenameConfig))
		})
	}
}

func TestInitFileConnections(t *testing.T) {
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
			require.Equal(t, tt.wantErr, initFileConnections() != nil)

			direction := storage.GetAppDir()
			require.True(t, storage.Exists(direction, envconst.FilenameConnections))
		})
	}
}

func TestInitLoggerFromConfig(t *testing.T) {
	type args struct {
		loggerType    string
		wantSetLogger logger.Logger
	}
	tests := []struct {
		name    string
		wantErr bool
		args    args
	}{
		{
			name:    "success set default logger",
			wantErr: false,
			args: args{
				loggerType:    "",
				wantSetLogger: logger.NewStorageLogger(),
			},
		},
		{
			name:    "success set console logger",
			wantErr: false,
			args: args{
				loggerType:    envconst.TypeConsoleLogger,
				wantSetLogger: logger.NewConsoleLogger(),
			},
		},
		{
			name:    "success set storage logger",
			wantErr: false,
			args: args{
				loggerType:    envconst.TypeStorageLogger,
				wantSetLogger: logger.NewStorageLogger(),
			},
		},
		{
			name: "success set combined logger",
			args: args{
				loggerType: envconst.TypeCombinedLogger,
				wantSetLogger: logger.NewCombinedLogger(
					logger.NewConsoleLogger(),
					logger.NewStorageLogger(),
				),
			},
			wantErr: false,
		},
		{
			name: "bad arg logger in configs",
			args: args{
				loggerType:    "badTypeLogger",
				wantSetLogger: logger.NewConsoleLogger(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.loggerType != "" {
				require.NoError(t, config.Set(envname.Logger, tt.args.loggerType))

				defer func() {
					require.NoError(t, config.Set(envname.Logger, envconst.TypeStorageLogger))
				}()
			}

			require.Equal(t, tt.wantErr, initLoggerFromConfig() != nil)

			if !tt.wantErr {
				require.Equal(t, reflect.TypeOf(logger.Get()), reflect.TypeOf(tt.args.wantSetLogger))
			}
		})
	}
}
