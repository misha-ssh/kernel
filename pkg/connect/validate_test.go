//go:build unit

package connect

import (
	"testing"
	"time"

	"github.com/misha-ssh/kernel/testutil"
	"github.com/stretchr/testify/require"
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
			if tt.wantErr {
				require.Error(t, tt.connection.Validate())
			} else {
				require.NoError(t, tt.connection.Validate())
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
			if tt.wantErr {
				require.Error(t, validateAlias(tt.alias))
			} else {
				require.NoError(t, validateAlias(tt.alias))
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
			if tt.wantErr {
				require.Error(t, validateLogin(tt.login))
			} else {
				require.NoError(t, validateLogin(tt.login))
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
			if tt.wantErr {
				require.Error(t, validateAddress(tt.address))
			} else {
				require.NoError(t, validateAddress(tt.address))
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name       string
		password   string
		privateKey string
		wantErr    bool
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
			name:       "success - skip validate pass",
			password:   "",
			privateKey: "path_to_pk_key",
			wantErr:    false,
		},
		{
			name:       "fail - none skip valida is empty key",
			password:   " ",
			privateKey: "",
			wantErr:    true,
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
			if tt.wantErr {
				require.Error(t, validatePassword(tt.password, tt.privateKey))
			} else {
				require.NoError(t, validatePassword(tt.password, tt.privateKey))
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
			if tt.wantErr {
				require.Error(t, validateDate(tt.date))
			} else {
				require.NoError(t, validateDate(tt.date))
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
			if tt.wantErr {
				require.Error(t, validateCreatedAt(tt.date))
			} else {
				require.NoError(t, validateCreatedAt(tt.date))
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
			if tt.wantErr {
				require.Error(t, validateUpdatedAt(tt.date))
			} else {
				require.NoError(t, validateUpdatedAt(tt.date))
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
			if tt.wantErr {
				require.Error(t, validatePort(tt.port))
			} else {
				require.NoError(t, validatePort(tt.port))
			}
		})
	}
}
