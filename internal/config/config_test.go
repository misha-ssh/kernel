package config

import (
	"testing"

	"github.com/misha-ssh/kernel/configs/envconst"
	"github.com/misha-ssh/kernel/internal/storage"
)

func TestStorage_Set(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		value   string
		wantErr bool
	}{
		{
			name:    "set numbers",
			key:     "test",
			value:   "123123",
			wantErr: false,
		},
		{
			name:    "set char early",
			key:     "testEarly",
			value:   "=",
			wantErr: true,
		},
		{
			name:    "set char early - rewrite key",
			key:     "test",
			value:   "=",
			wantErr: true,
		},
		{
			name:    "set numeric",
			key:     "randomKey",
			value:   "0",
			wantErr: false,
		},
		{
			name:    "set one char",
			key:     "test",
			value:   "z",
			wantErr: false,
		},
		{
			name:    "set letters",
			key:     "test",
			value:   "randomString",
			wantErr: false,
		},
		{
			name:    "set letters and numbers",
			key:     "test",
			value:   "randomString123123",
			wantErr: false,
		},
		{
			name:    "empty key",
			key:     "",
			value:   "test",
			wantErr: true,
		},
		{
			name:    "key is spaces",
			key:     "   ",
			value:   "test",
			wantErr: true,
		},
		{
			name:    "empty value",
			key:     "test",
			value:   "",
			wantErr: true,
		},
		{
			name:    "values is spaces",
			key:     "test",
			value:   "   ",
			wantErr: true,
		},
		{
			name:    "set key with spaces",
			key:     "test 2",
			value:   "2",
			wantErr: true,
		},
		{
			name:    "set value with spaces",
			key:     "test2",
			value:   "a a",
			wantErr: true,
		},
		{
			name:    "set value with random register",
			key:     "RaRd",
			value:   "test",
			wantErr: false,
		},
	}

	filename := envconst.FilenameConfig
	direction := storage.GetAppDir()

	if !storage.Exists(direction, filename) {
		err := storage.Create(direction, filename)
		if err != nil {
			t.Fatal(err)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Set(tt.key, tt.value)

			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := Get(tt.key)

			if (got != tt.value) != tt.wantErr {
				t.Errorf("got: %v != want: %v", got, tt.wantErr)
			}
		})
	}
}

func TestStorage_Get(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		want       string
		isSetValue bool
	}{
		{
			name:       "get existing value",
			key:        "test",
			want:       "testData",
			isSetValue: true,
		},
		{
			name:       "get rewriting value",
			key:        "test",
			want:       "rewriteData",
			isSetValue: true,
		},
		{
			name:       "get numeric value",
			key:        "newTestKey",
			want:       "0",
			isSetValue: true,
		},
		{
			name:       "get lat set value",
			key:        "test",
			want:       "rewriteData",
			isSetValue: false,
		},
		{
			name:       "get non existing value",
			key:        "empty",
			want:       "",
			isSetValue: false,
		},
		{
			name:       "get empty key",
			key:        "",
			want:       "",
			isSetValue: false,
		},
	}

	filename := envconst.FilenameConfig
	direction := storage.GetAppDir()

	if !storage.Exists(direction, filename) {
		err := storage.Create(direction, filename)
		if err != nil {
			t.Fatal(err)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isSetValue {
				err := Set(tt.key, tt.want)
				if err != nil {
					t.Errorf("Set() error = %v", err)
				}
			}

			got := Get(tt.key)
			if got != tt.want {
				t.Errorf("got: %v != want: %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Exists(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		isCreateKey bool
		want        bool
	}{
		{
			name:        "key exists",
			key:         "test",
			isCreateKey: true,
			want:        true,
		},
		{
			name:        "key exists with random register",
			key:         "tEsT",
			isCreateKey: false,
			want:        true,
		},
		{
			name:        "key dont exists",
			key:         "dontExistKey",
			isCreateKey: false,
			want:        false,
		},
		{
			name:        "empty key",
			key:         "",
			isCreateKey: false,
			want:        false,
		},
	}

	filename := envconst.FilenameConfig
	direction := storage.GetAppDir()

	if !storage.Exists(direction, filename) {
		err := storage.Create(direction, filename)
		if err != nil {
			t.Fatal(err)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isCreateKey {
				err := Set(tt.key, "test")
				if err != nil {
					t.Errorf("Set() error = %v", err)
				}
			}

			got := Exists(tt.key)
			if got != tt.want {
				t.Errorf("got: %v != want: %v", got, tt.want)
			}
		})
	}
}
