package connect

import (
	"testing"
	"time"

	"github.com/misha-ssh/kernel/testutil"
)

func TestConnectValidate(t *testing.T) {
	tests := []struct {
		name       string
		connection *Connect
		wantErr    bool
	}{
		{
			name: "success - all fields valid",
			connection: &Connect{
				Alias:     testutil.RandomString(),
				Login:     "username",
				Address:   "example.com",
				Password:  "password123",
				CreatedAt: time.Now().Add(-time.Hour).Format(time.RFC3339),
				UpdatedAt: time.Now().Add(-30 * time.Minute).Format(time.RFC3339),
				SshOptions: &SshOptions{
					Port: 22,
				},
			},
			wantErr: false,
		},
		{
			name: "fail - invalid alias",
			connection: &Connect{
				Alias:     "invalid@alias",
				Login:     "username",
				Address:   "example.com",
				Password:  "password123",
				CreatedAt: time.Now().Format(time.RFC3339),
				UpdatedAt: time.Now().Format(time.RFC3339),
				SshOptions: &SshOptions{
					Port: 22,
				},
			},
			wantErr: true,
		},
		{
			name: "fail - invalid port",
			connection: &Connect{
				Alias:     testutil.RandomString(),
				Login:     "username",
				Address:   "example.com",
				Password:  "password123",
				CreatedAt: time.Now().Format(time.RFC3339),
				UpdatedAt: time.Now().Format(time.RFC3339),
				SshOptions: &SshOptions{
					Port: 70000,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.connection.Validate(); (err != nil) != tt.wantErr {
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

func TestValidateLogin(t *testing.T) {
	tests := []struct {
		name    string
		login   string
		wantErr bool
	}{
		{
			name:    "success - valid login",
			login:   "username",
			wantErr: false,
		},
		{
			name:    "success - with dots and dashes",
			login:   "user.name-123",
			wantErr: false,
		},
		{
			name:    "success - with underscore",
			login:   "user_name",
			wantErr: false,
		},
		{
			name:    "fail - empty login",
			login:   "",
			wantErr: true,
		},
		{
			name:    "fail - spaces only",
			login:   "   ",
			wantErr: true,
		},
		{
			name:    "fail - too long",
			login:   "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			wantErr: true,
		},
		{
			name:    "fail - russian chars",
			login:   "пользователь",
			wantErr: true,
		},
		{
			name:    "fail - special chars",
			login:   "user@name",
			wantErr: true,
		},
		{
			name:    "fail - with spaces",
			login:   "user name",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateLogin(tt.login); (err != nil) != tt.wantErr {
				t.Errorf("validateLogin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		wantErr bool
	}{
		{
			name:    "success - valid domain",
			address: "example.com",
			wantErr: false,
		},
		{
			name:    "success - valid IP address",
			address: "192.168.1.1",
			wantErr: false,
		},
		{
			name:    "success - IPv6 address",
			address: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			wantErr: false,
		},
		{
			name:    "success - subdomain",
			address: "sub.example.com",
			wantErr: false,
		},
		{
			name:    "success - with numbers",
			address: "server-123.example.com",
			wantErr: false,
		},
		{
			name:    "success - russian real domain",
			address: "xn--80acbbp8bh.xn--p1ai",
			wantErr: false,
		},
		{
			name:    "fail - empty address",
			address: "",
			wantErr: true,
		},
		{
			name:    "fail - spaces only",
			address: "   ",
			wantErr: true,
		},
		{
			name:    "fail - invalid chars",
			address: "example.com/path",
			wantErr: true,
		},
		{
			name:    "fail - with underscore",
			address: "server_name.com",
			wantErr: true,
		},
		{
			name:    "fail - long domain",
			address: "exampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexampleexample.com",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateAddress(tt.address); (err != nil) != tt.wantErr {
				t.Errorf("validateAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "success - valid password",
			password: "password123",
			wantErr:  false,
		},
		{
			name:     "success - minimum length",
			password: "pass",
			wantErr:  false,
		},
		{
			name:     "success - special chars",
			password: "p@ssw0rd!",
			wantErr:  false,
		},
		{
			name:     "fail - empty password",
			password: "",
			wantErr:  true,
		},
		{
			name:     "fail - spaces only",
			password: "    ",
			wantErr:  true,
		},
		{
			name:     "fail - too short",
			password: "abc",
			wantErr:  true,
		},
		{
			name:     "fail - too long",
			password: "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePassword(tt.password); (err != nil) != tt.wantErr {
				t.Errorf("validatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateDate(t *testing.T) {
	now := time.Now().UTC()
	validPastDate := now.Add(-time.Hour).Format(time.RFC3339)
	validFutureDate := now.Add(time.Hour).Format(time.RFC3339)
	invalidFormat := "2023-01-01 12:00:00"

	tests := []struct {
		name    string
		date    string
		wantErr bool
	}{
		{
			name:    "success - valid past date",
			date:    validPastDate,
			wantErr: false,
		},
		{
			name:    "success - valid current date",
			date:    now.Format(time.RFC3339),
			wantErr: false,
		},
		{
			name:    "fail - empty date",
			date:    "",
			wantErr: true,
		},
		{
			name:    "fail - spaces only",
			date:    "   ",
			wantErr: true,
		},
		{
			name:    "fail - future date",
			date:    validFutureDate,
			wantErr: true,
		},
		{
			name:    "fail - invalid format",
			date:    invalidFormat,
			wantErr: true,
		},
		{
			name:    "fail - malformed date",
			date:    "2023-13-45T25:61:61Z",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateDate(tt.date); (err != nil) != tt.wantErr {
				t.Errorf("validateDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCreatedAt(t *testing.T) {
	validDate := time.Now().Add(-time.Hour).Format(time.RFC3339)

	tests := []struct {
		name    string
		date    string
		wantErr bool
	}{
		{
			name:    "success - valid date",
			date:    validDate,
			wantErr: false,
		},
		{
			name:    "fail - empty date",
			date:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreatedAt(tt.date); (err != nil) != tt.wantErr {
				t.Errorf("validateCreatedAt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateUpdatedAt(t *testing.T) {
	validDate := time.Now().Add(-time.Hour).Format(time.RFC3339)

	tests := []struct {
		name    string
		date    string
		wantErr bool
	}{
		{
			name:    "success - valid date",
			date:    validDate,
			wantErr: false,
		},
		{
			name:    "fail - empty date",
			date:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateUpdatedAt(tt.date); (err != nil) != tt.wantErr {
				t.Errorf("validateUpdatedAt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePort(t *testing.T) {
	tests := []struct {
		name    string
		port    int
		wantErr bool
	}{
		{
			name:    "success - valid port",
			port:    8080,
			wantErr: false,
		},
		{
			name:    "success - minimum port",
			port:    1,
			wantErr: false,
		},
		{
			name:    "success - maximum port",
			port:    65535,
			wantErr: false,
		},
		{
			name:    "fail - port too small",
			port:    0,
			wantErr: true,
		},
		{
			name:    "fail - port too large",
			port:    65536,
			wantErr: true,
		},
		{
			name:    "fail - negative port",
			port:    -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePort(tt.port); (err != nil) != tt.wantErr {
				t.Errorf("validatePort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
