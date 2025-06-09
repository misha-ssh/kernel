package logger

import (
	"math/rand"
	"testing"
)

func TestCombinedLogger_Debug(t *testing.T) {
	type fields struct {
		loggers []Logger
	}
	type args struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "console logger",
			fields: fields{
				loggers: []Logger{
					NewConsoleLogger(),
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "storage logger",
			fields: fields{
				loggers: []Logger{
					NewStorageLogger(),
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "console + storage logger",
			fields: fields{
				loggers: []Logger{
					NewConsoleLogger(),
					NewStorageLogger(),
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := NewCombinedLogger(tt.fields.loggers...)

			defer func() {
				if r := recover(); r != nil {
					t.Error("Debug() is panicked")
				}
			}()

			cl.Debug(tt.args.value)
		})
	}
}

func TestCombinedLogger_Error(t *testing.T) {
	type fields struct {
		loggers []Logger
	}
	type args struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "console logger",
			fields: fields{
				loggers: []Logger{
					NewConsoleLogger(),
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "storage logger",
			fields: fields{
				loggers: []Logger{
					NewStorageLogger(),
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "console + storage logger",
			fields: fields{
				loggers: []Logger{
					NewConsoleLogger(),
					NewStorageLogger(),
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := NewCombinedLogger(tt.fields.loggers...)

			defer func() {
				if r := recover(); r != nil {
					t.Error("Error() is panicked")
				}
			}()

			cl.Error(tt.args.value)
		})
	}
}

func TestCombinedLogger_Info(t *testing.T) {
	type fields struct {
		loggers []Logger
	}
	type args struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "console logger",
			fields: fields{
				loggers: []Logger{
					NewConsoleLogger(),
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "storage logger",
			fields: fields{
				loggers: []Logger{
					NewStorageLogger(),
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "console + storage logger",
			fields: fields{
				loggers: []Logger{
					NewConsoleLogger(),
					NewStorageLogger(),
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := NewCombinedLogger(tt.fields.loggers...)

			defer func() {
				if r := recover(); r != nil {
					t.Error("Info() is panicked")
				}
			}()

			cl.Info(tt.args.value)
		})
	}
}

func TestCombinedLogger_Warn(t *testing.T) {
	type fields struct {
		loggers []Logger
	}
	type args struct {
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "console logger",
			fields: fields{
				loggers: []Logger{
					NewConsoleLogger(),
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "storage logger",
			fields: fields{
				loggers: []Logger{
					NewStorageLogger(),
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "console + storage logger",
			fields: fields{
				loggers: []Logger{
					NewConsoleLogger(),
					NewStorageLogger(),
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := NewCombinedLogger(tt.fields.loggers...)

			defer func() {
				if r := recover(); r != nil {
					t.Error("Warn() is panicked")
				}
			}()

			cl.Warn(tt.args.value)
		})
	}
}
