package logger

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorageLogger_Error(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				Error(tt.value)
			})
		})
	}
}

func TestStorageLogger_Warn(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				Warn(tt.value)
			})
		})
	}
}

func TestStorageLogger_Info(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				Info(tt.value)
			})
		})
	}
}

func TestStorageLogger_Debug(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				Debug(tt.value)
			})
		})
	}
}

func TestStorageLogger_log(t *testing.T) {
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

	storageLogger := New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storageLogger.log(tt.value, tt.status)

			assert.NoError(t, err)
		})
	}
}
