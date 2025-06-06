package config

import (
	"testing"
)

func Test_Init(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success init",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_initCryptKey(t *testing.T) {
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
			if err := initCryptKey(); (err != nil) != tt.wantErr {
				t.Errorf("initCryptKey() error = %v, wantErr %v", err, tt.wantErr)
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
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := initFileConfig(); (err != nil) != tt.wantErr {
				t.Errorf("initFileConfig() error = %v, wantErr %v", err, tt.wantErr)
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
		})
	}
}
