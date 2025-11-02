//go:build unit

package logger

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
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
				require.Nil(t, recover())
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
				require.Nil(t, recover())
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
				require.Nil(t, recover())
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
				require.Nil(t, recover())
			}()

			Info(tt.value)
		})
	}
}
