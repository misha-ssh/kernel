package connect

import (
	"reflect"
	"testing"

	"golang.org/x/crypto/ssh"
)

func TestSshConnect(t *testing.T) {
	type args struct {
		session *ssh.Session
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "nil session",
			args:    args{session: nil},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Ssh{}
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantErr {
						t.Errorf("Connect() unexpected panic = %v", r)
					}
				}
			}()

			err := s.Connect(tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSshSession(t *testing.T) {
	type args struct {
		connection *Connect
	}
	tests := []struct {
		name    string
		args    args
		want    *ssh.Session
		wantErr bool
	}{
		{
			name: "empty connection data",
			args: args{connection: &Connect{
				SshOptions: &SshOptions{},
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid connection data",
			args: args{connection: &Connect{
				Address:  "invalid-host",
				Login:    "user",
				Password: "pass",
				SshOptions: &SshOptions{
					Port: 22,
				},
			}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Ssh{}
			got, err := s.Session(tt.args.connection)
			if (err != nil) != tt.wantErr {
				t.Errorf("Session() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuth(t *testing.T) {
	type args struct {
		connection *Connect
	}
	tests := []struct {
		name    string
		args    args
		want    []ssh.AuthMethod
		wantErr bool
	}{
		{
			name: "password authentication",
			args: args{connection: &Connect{
				Password: "testpassword",
				SshOptions: &SshOptions{
					PrivateKey: "",
				},
			}},
			want:    []ssh.AuthMethod{ssh.Password("testpassword")},
			wantErr: false,
		},
		{
			name: "private key authentication with empty key",
			args: args{connection: &Connect{
				SshOptions: &SshOptions{
					PrivateKey: "",
				},
			}},
			want:    []ssh.AuthMethod{ssh.Password("")},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := auth(tt.args.connection)
			if (err != nil) != tt.wantErr {
				t.Errorf("auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("auth() got length = %v, want length %v", len(got), len(tt.want))
				return
			}
		})
	}
}

func TestCreateTerminalSession(t *testing.T) {
	type args struct {
		session *ssh.Session
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "nil session",
			args:    args{session: nil},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantErr {
						t.Errorf("createTerminalSession() unexpected panic = %v", r)
					}
				}
			}()

			err := createTerminalSession(tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("createTerminalSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetClient(t *testing.T) {
	type args struct {
		hostWithPort string
		config       *ssh.ClientConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *ssh.Client
		wantErr bool
	}{
		{
			name: "invalid host",
			args: args{
				hostWithPort: "invalid-host:22",
				config: &ssh.ClientConfig{
					Timeout:         Timeout,
					User:            "test",
					Auth:            []ssh.AuthMethod{ssh.Password("test")},
					HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty host",
			args: args{
				hostWithPort: "",
				config: &ssh.ClientConfig{
					Timeout: Timeout,
					User:    "test",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getClient(tt.args.hostWithPort, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("getClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetClientConfig(t *testing.T) {
	type args struct {
		connection *Connect
	}
	tests := []struct {
		name    string
		args    args
		want    *ssh.ClientConfig
		wantErr bool
	}{
		{
			name: "valid connection with password",
			args: args{connection: &Connect{
				Login:      "testuser",
				Password:   "testpass",
				SshOptions: &SshOptions{},
			}},
			want: &ssh.ClientConfig{
				Timeout:         Timeout,
				User:            "testuser",
				Auth:            []ssh.AuthMethod{ssh.Password("testpass")},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			},
			wantErr: false,
		},
		{
			name: "invalid connection with password",
			args: args{connection: &Connect{
				Login:      "",
				Password:   "",
				SshOptions: &SshOptions{},
			}},
			want: &ssh.ClientConfig{
				Timeout:         Timeout,
				User:            "",
				Auth:            []ssh.AuthMethod{ssh.Password("")},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getClientConfig(tt.args.connection)
			if (err != nil) != tt.wantErr {
				t.Errorf("getClientConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Timeout != tt.want.Timeout {
				t.Errorf("getClientConfig() timeout = %v, want %v", got.Timeout, tt.want.Timeout)
			}
			if got.User != tt.want.User {
				t.Errorf("getClientConfig() user = %v, want %v", got.User, tt.want.User)
			}
			if len(got.Auth) != len(tt.want.Auth) {
				t.Errorf("getClientConfig() auth methods count = %v, want %v", len(got.Auth), len(tt.want.Auth))
			}
		})
	}
}

func TestGetSession(t *testing.T) {
	type args struct {
		client *ssh.Client
	}
	tests := []struct {
		name    string
		args    args
		want    *ssh.Session
		wantErr bool
	}{
		{
			name:    "nil client",
			args:    args{client: nil},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantErr {
						t.Errorf("getSession() unexpected panic = %v", r)
					}
				}
			}()

			got, err := getSession(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSession() got = %v, want %v", got, tt.want)
			}
		})
	}
}
