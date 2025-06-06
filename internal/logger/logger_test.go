package logger

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
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

	consoleLogger := ConsoleLogger{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				consoleLogger.Error(tt.value)
			})
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

	consoleLogger := ConsoleLogger{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				consoleLogger.Debug(tt.value)
			})
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

	consoleLogger := ConsoleLogger{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				consoleLogger.Warn(tt.value)
			})
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

	consoleLogger := ConsoleLogger{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				consoleLogger.Info(tt.value)
			})
		})
	}
}
