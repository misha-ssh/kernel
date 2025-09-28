//go:build unit

package logger

import (
	"math/rand"
	"testing"
)

func TestLogger_Error(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Error() is panicked")
				}
			}()

			Error(tt.value)
		})
	}
}

func TestLogger_Debug(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Debug() is panicked")
				}
			}()

			Debug(tt.value)
		})
	}
}

func TestLogger_Warn(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Warn() is panicked")
				}
			}()

			Warn(tt.value)
		})
	}
}

func TestLogger_Info(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error("Info() is panicked")
				}
			}()

			Info(tt.value)
		})
	}
}
