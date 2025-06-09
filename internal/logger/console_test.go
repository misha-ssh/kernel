package logger

import (
	"math/rand"
	"testing"
)

func TestConsoleLogger_Error(t *testing.T) {
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

	consoleLogger := NewConsoleLogger()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Error() is panicked")
				}
			}()

			consoleLogger.Error(tt.value)
		})
	}
}

func TestConsoleLogger_Warn(t *testing.T) {
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

	consoleLogger := NewConsoleLogger()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Warn() is panicked")
				}
			}()

			consoleLogger.Warn(tt.value)
		})
	}
}

func TestConsoleLogger_Info(t *testing.T) {
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

	consoleLogger := NewConsoleLogger()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Info() is panicked")
				}
			}()

			consoleLogger.Info(tt.value)
		})
	}
}

func TestConsoleLogger_Debug(t *testing.T) {
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

	consoleLogger := NewConsoleLogger()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Debug() is panicked")
				}
			}()

			consoleLogger.Debug(tt.value)
		})
	}
}

func TestConsoleLogger_log(t *testing.T) {
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

	consoleLogger := NewConsoleLogger()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := consoleLogger.log(tt.value, tt.status)
			if err != nil {
				t.Errorf("log error: %v", err)
			}
		})
	}
}
