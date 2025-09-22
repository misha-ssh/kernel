package connect

import (
	"testing"
)

func TestConnectValidate(t *testing.T) {
	type args struct {
		connection *Connect
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - simple chars",
			args: args{
				connection: &Connect{
					Alias: "aasd",
				},
			},
			wantErr: false,
		},
		{
			name: "success - numbers",
			args: args{
				connection: &Connect{
					Alias: "123123",
				},
			},
			wantErr: false,
		},
		{
			name: "fail - ru chars",
			args: args{
				connection: &Connect{
					Alias: "йцуйцу",
				},
			},
			wantErr: true,
		},
		{
			name: "fail - empty alias",
			args: args{
				connection: &Connect{
					Alias: "",
				},
			},
			wantErr: true,
		},
		{
			name: "fail - spaces",
			args: args{
				connection: &Connect{
					Alias: "  ",
				},
			},
			wantErr: true,
		},
		{
			name: "fail - with /",
			args: args{
				connection: &Connect{
					Alias: "alias/",
				},
			},
			wantErr: true,
		},
		{
			name: "fail - with $",
			args: args{
				connection: &Connect{
					Alias: "$test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.args.connection.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAlias(t *testing.T) {
	tests := []struct {
		name    string
		alias   string
		wantErr bool
	}{
		{
			name:    "success - simple chars",
			alias:   "aasd",
			wantErr: false,
		},
		{
			name:    "success - numbers",
			alias:   "123123",
			wantErr: false,
		},
		{
			name:    "fail - ru chars",
			alias:   "ыфвдыджфлв",
			wantErr: true,
		},
		{
			name:    "fail - empty alias",
			alias:   "",
			wantErr: true,
		},
		{
			name:    "fail - spaces",
			alias:   "  ",
			wantErr: true,
		},
		{
			name:    "fail - with /",
			alias:   "alias/",
			wantErr: true,
		},
		{
			name:    "fail - with $",
			alias:   "$test",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateAlias(tt.alias); (err != nil) != tt.wantErr {
				t.Errorf("validateAlias() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
