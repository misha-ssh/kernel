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
	"github.com/zalando/go-keyring"
)

func Test_Init(t *testing.T) {
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
				if r := recover(); r != nil {
					t.Error("Init() is panicked")
				}
			}()

			Init()
		})
	}
}

func Test_initCryptKey(t *testing.T) {
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
			if err := initCryptKey(); (err != nil) != tt.wantErr {
				t.Errorf("initCryptKey() error = %v, wantErr %v", err, tt.wantErr)
			}

			currentUser, err := user.Current()
			if err != nil {
				t.Errorf("initCryptKey() error = %v", err)
			}

			username := currentUser.Username
			service := envconst.NameServiceCryptKey
			cryptKey, _ := keyring.Get(service, username)

			if len(cryptKey) != crypto.SizeKey {
				t.Errorf("initCryptKey() error = %v, CryptKey size is %v", err, crypto.SizeKey)
			}
		})
	}
}

func Test_initFileConfig(t *testing.T) {
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
			if err := initFileConfig(); (err != nil) != tt.wantErr {
				t.Errorf("initFileConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			direction := storage.GetAppDir()

			if !storage.Exists(direction, envconst.FilenameConfig) {
				t.Error("initFileConfig() dont create file")
			}
		})
	}
}

func Test_initFileConnections(t *testing.T) {
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
			if err := initFileConnections(); (err != nil) != tt.wantErr {
				t.Errorf("initFileConnections() error = %v, wantErr %v", err, tt.wantErr)
			}

			direction := storage.GetAppDir()

			if !storage.Exists(direction, envconst.FilenameConnections) {
				t.Error("initFileConnections() dont create file")
			}
		})
	}
}

func Test_initLoggerFromConfig(t *testing.T) {
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
				err := config.Set(envname.Logger, tt.args.loggerType)
				if err != nil {
					t.Errorf("initLoggerFromConfig() error = %v", err)
				}

				// set default type logger before completed test
				defer func() {
					if err = config.Set(envname.Logger, envconst.TypeStorageLogger); err != nil {
						t.Errorf("Set() error = %v", err)
					}
				}()
			}

			if err := initLoggerFromConfig(); (err != nil) != tt.wantErr {
				t.Errorf("initLoggerFromConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				loggerSetting := logger.Get()

				if reflect.TypeOf(loggerSetting).String() != reflect.TypeOf(tt.args.wantSetLogger).String() {
					t.Errorf("logger from configs: %v != %v", loggerSetting, tt.args.wantSetLogger)
				}
			}
		})
	}
}
