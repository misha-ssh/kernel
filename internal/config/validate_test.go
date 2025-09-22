package config

import (
	"testing"
)

func TestValidateKey(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		wantErr bool
	}{
		{
			name:    "success - validate",
			key:     "test",
			wantErr: false,
		},
		{
			name:    "fail - key is numbers",
			key:     "123123123",
			wantErr: true,
		},
		{
			name:    "fail - key is ru chars",
			key:     "тест",
			wantErr: true,
		},
		{
			name:    "fail - key is empty",
			key:     "",
			wantErr: true,
		},
		{
			name:    "fail - chars with spaces",
			key:     "test  two",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateKey(tt.key); (err != nil) != tt.wantErr {
				t.Errorf("validateKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOnEmptyString(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{
			name:    "success - validate",
			value:   "not_empty_string",
			wantErr: false,
		},
		{
			name:    "success - validate",
			value:   "space string",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateOnEmptyString(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("validateOnEmptyString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateValue(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "valid alphanumeric value",
			args:    args{value: "abc123"},
			wantErr: false,
		},
		{
			name:    "valid value with dots",
			args:    args{value: "abc.123"},
			wantErr: false,
		},
		{
			name:    "valid value with hyphens",
			args:    args{value: "abc-123"},
			wantErr: false,
		},
		{
			name:    "valid value with underscores",
			args:    args{value: "abc_123"},
			wantErr: false,
		},
		{
			name:    "valid value with mixed special characters",
			args:    args{value: "a.b-c_d"},
			wantErr: false,
		},
		{
			name:    "valid uppercase letters",
			args:    args{value: "ABC123"},
			wantErr: false,
		},
		{
			name:    "valid mixed case",
			args:    args{value: "AbC123"},
			wantErr: false,
		},
		{
			name:    "empty string - should fail",
			args:    args{value: ""},
			wantErr: true,
		},
		{
			name:    "value with spaces - should fail",
			args:    args{value: "abc 123"},
			wantErr: true,
		},
		{
			name:    "value with special characters - should fail",
			args:    args{value: "abc@123"},
			wantErr: true,
		},
		{
			name:    "value with slashes - should fail",
			args:    args{value: "abc/123"},
			wantErr: true,
		},
		{
			name:    "value with parentheses - should fail",
			args:    args{value: "abc(123)"},
			wantErr: true,
		},
		{
			name:    "value with unicode characters - should fail",
			args:    args{value: "abcтест123"},
			wantErr: true,
		},
		{
			name:    "value with emoji - should fail",
			args:    args{value: "abc😊123"},
			wantErr: true,
		},
		{
			name:    "only dots and hyphens - valid",
			args:    args{value: ".-_"},
			wantErr: false,
		},
		{
			name:    "starts with special character - valid",
			args:    args{value: ".abc123"},
			wantErr: false,
		},
		{
			name:    "ends with special character - valid",
			args:    args{value: "abc123."},
			wantErr: false,
		},
		{
			name:    "long valid value",
			args:    args{value: "a1.b2-c3_d4.e5-f6_g7.h8-i9_j0"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateValue(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("validateValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
