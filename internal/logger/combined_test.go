package logger

import (
	"math/rand"
	"testing"

	"github.com/ssh-connection-manager/kernel/v2/internal/storage"
	"github.com/stretchr/testify/assert"
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
					&ConsoleLogger{},
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "storage logger(local)",
			fields: fields{
				loggers: []Logger{
					&StorageLogger{
						Storage: &storage.FileStorage{
							Direction: t.TempDir(),
						},
					},
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "console logger + storage logger(local)",
			fields: fields{
				loggers: []Logger{
					&ConsoleLogger{},
					&StorageLogger{
						Storage: &storage.FileStorage{
							Direction: t.TempDir(),
						},
					},
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &CombinedLogger{
				loggers: tt.fields.loggers,
			}

			assert.NotPanics(t, func() {
				cl.Debug(tt.args.value)
			})
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
					&ConsoleLogger{},
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "storage logger(local)",
			fields: fields{
				loggers: []Logger{
					&StorageLogger{
						Storage: &storage.FileStorage{
							Direction: t.TempDir(),
						},
					},
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "console logger + storage logger(local)",
			fields: fields{
				loggers: []Logger{
					&ConsoleLogger{},
					&StorageLogger{
						Storage: &storage.FileStorage{
							Direction: t.TempDir(),
						},
					},
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &CombinedLogger{
				loggers: tt.fields.loggers,
			}

			assert.NotPanics(t, func() {
				cl.Error(tt.args.value)
			})
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
					&ConsoleLogger{},
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "storage logger(local)",
			fields: fields{
				loggers: []Logger{
					&StorageLogger{
						Storage: &storage.FileStorage{
							Direction: t.TempDir(),
						},
					},
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "console logger + storage logger(local)",
			fields: fields{
				loggers: []Logger{
					&ConsoleLogger{},
					&StorageLogger{
						Storage: &storage.FileStorage{
							Direction: t.TempDir(),
						},
					},
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &CombinedLogger{
				loggers: tt.fields.loggers,
			}

			assert.NotPanics(t, func() {
				cl.Info(tt.args.value)
			})
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
					&ConsoleLogger{},
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "storage logger(local)",
			fields: fields{
				loggers: []Logger{
					&StorageLogger{
						Storage: &storage.FileStorage{
							Direction: t.TempDir(),
						},
					},
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
		{
			name: "console logger + storage logger(local)",
			fields: fields{
				loggers: []Logger{
					&ConsoleLogger{},
					&StorageLogger{
						Storage: &storage.FileStorage{
							Direction: t.TempDir(),
						},
					},
				},
			},
			args: args{
				value: rand.Int(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &CombinedLogger{
				loggers: tt.fields.loggers,
			}

			assert.NotPanics(t, func() {
				cl.Warn(tt.args.value)
			})
		})
	}
}
