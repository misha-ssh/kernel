//go:build unit

package logger

import (
	"math/rand"
	"testing"
)

func TestStorageLogger_Error(t *testing.T) {
	tests := []struct {
		name   string
		status StatusLog
		value  any
	}{
		{
			name:  "success - rand int",
			value: rand.Int(),
		},
		{
			name:  "success - default string",
			value: "test",
		},
	}

	storageLogger := NewStorageLogger()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Error() is panicked")
				}
			}()

			storageLogger.Error(tt.value)
		})
	}
}

func TestStorageLogger_Warn(t *testing.T) {
	tests := []struct {
		name   string
		status StatusLog
		value  any
	}{
		{
			name:  "success - rand int",
			value: rand.Int(),
		},
		{
			name:  "success - default string",
			value: "test",
		},
	}

	storageLogger := NewStorageLogger()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Warn() is panicked")
				}
			}()

			storageLogger.Warn(tt.value)
		})
	}
}

func TestStorageLogger_Info(t *testing.T) {
	tests := []struct {
		name   string
		status StatusLog
		value  any
	}{
		{
			name:  "success - rand int",
			value: rand.Int(),
		},
		{
			name:  "success - default string",
			value: "test",
		},
	}

	storageLogger := NewStorageLogger()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Info() is panicked")
				}
			}()

			storageLogger.Info(tt.value)
		})
	}
}

func TestStorageLogger_Debug(t *testing.T) {
	tests := []struct {
		name   string
		status StatusLog
		value  any
	}{
		{
			name:  "success - rand int",
			value: rand.Int(),
		},
		{
			name:  "success - default string",
			value: "test",
		},
	}

	storageLogger := NewStorageLogger()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Debug() is panicked")
				}
			}()

			storageLogger.Debug(tt.value)
		})
	}
}

func TestStorageLogger_log(t *testing.T) {
	tests := []struct {
		name   string
		status StatusLog
		value  any
	}{
		{
			name:   "success - set value with info status",
			status: InfoStatus,
			value:  rand.Int(),
		},
		{
			name:   "success - set value with error status",
			status: ErrorStatus,
			value:  rand.Int(),
		},
		{
			name:   "success - set value with debug status",
			status: DebugStatus,
			value:  rand.Int(),
		},
		{
			name:   "success - set value with warn status",
			status: WarnStatus,
			value:  rand.Int(),
		},
	}

	storageLogger := NewStorageLogger()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storageLogger.log(tt.value, tt.status)
			if err != nil {
				t.Errorf("log() error = %v", err)
			}
		})
	}
}
