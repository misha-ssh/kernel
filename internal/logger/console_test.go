package logger

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsoleLogger_Error(t *testing.T) {
	tests := []struct {
		name    string
		status  Status
		isPanic bool
		value   any
	}{
		{
			name:    "success - rand int",
			isPanic: false,
			value:   rand.Int(),
		},
		{
			name:    "success - default string",
			isPanic: false,
			value:   "test",
		},
	}

	consoleLogger := ConsoleLogger{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				consoleLogger.Error(tt.value)
			})
		})
	}
}

func TestConsoleLogger_Warn(t *testing.T) {
	tests := []struct {
		name    string
		status  Status
		isPanic bool
		value   any
	}{
		{
			name:    "success - rand int",
			isPanic: false,
			value:   rand.Int(),
		},
		{
			name:    "success - default string",
			isPanic: false,
			value:   "test",
		},
	}

	consoleLogger := ConsoleLogger{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				consoleLogger.Warn(tt.value)
			})
		})
	}
}

func TestConsoleLogger_Info(t *testing.T) {
	tests := []struct {
		name    string
		status  Status
		isPanic bool
		value   any
	}{
		{
			name:    "success - rand int",
			isPanic: false,
			value:   rand.Int(),
		},
		{
			name:    "success - default string",
			isPanic: false,
			value:   "test",
		},
	}

	consoleLogger := ConsoleLogger{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				consoleLogger.Info(tt.value)
			})
		})
	}
}

func TestConsoleLogger_Debug(t *testing.T) {
	tests := []struct {
		name    string
		status  Status
		isPanic bool
		value   any
	}{
		{
			name:    "success - rand int",
			isPanic: false,
			value:   rand.Int(),
		},
		{
			name:    "success - default string",
			isPanic: false,
			value:   "test",
		},
	}

	consoleLogger := ConsoleLogger{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				consoleLogger.Debug(tt.value)
			})
		})
	}
}

func TestConsoleLogger_log(t *testing.T) {
	tests := []struct {
		name   string
		status Status
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

	consoleLogger := ConsoleLogger{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := consoleLogger.log(tt.value, tt.status)
			assert.NoError(t, err)
		})
	}
}
