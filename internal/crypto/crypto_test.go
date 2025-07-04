package crypto

import "testing"

func TestStorage_Decrypt(t *testing.T) {
	type args struct {
		plaintext string
		key       string
	}
	tests := []struct {
		name        string
		args        args
		want        string
		wantErr     bool
		generateKey bool
	}{
		{
			name: "success",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:        "hello world",
			wantErr:     false,
			generateKey: true,
		},
		{
			name: "empty (string) plaintext",
			args: args{
				plaintext: "",
				key:       "",
			},
			want:        "",
			wantErr:     false,
			generateKey: true,
		},
		{
			name: "empty key",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:        "hello world",
			wantErr:     true,
			generateKey: false,
		},
		{
			name: "long size key",
			args: args{
				plaintext: "hello world",
				key:       "32-byte-long-key-1234567890123456",
			},
			want:        "hello world",
			wantErr:     true,
			generateKey: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			if tt.generateKey {
				tt.args.key, err = GenerateKey()
				if err != nil {
					t.Errorf("GenerateKey() error = %v", err)
				}
			}

			cryptText, err := Encrypt(tt.args.plaintext, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := Decrypt(cryptText, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
			}

			if (got != tt.want) != tt.wantErr {
				t.Errorf("got: %v != want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Encrypt(t *testing.T) {
	type args struct {
		plaintext string
		key       string
	}
	tests := []struct {
		name        string
		args        args
		want        string
		wantErr     bool
		generateKey bool
	}{
		{
			name: "success",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:        "hello world",
			wantErr:     false,
			generateKey: true,
		},
		{
			name: "empty (string) plaintext",
			args: args{
				plaintext: "",
				key:       "",
			},
			want:        "",
			wantErr:     false,
			generateKey: true,
		},
		{
			name: "empty key",
			args: args{
				plaintext: "hello world",
				key:       "",
			},
			want:        "hello world",
			wantErr:     true,
			generateKey: false,
		},
		{
			name: "long size key",
			args: args{
				plaintext: "hello world",
				key:       "32-byte-long-key-1234567890123456",
			},
			want:        "hello world",
			wantErr:     true,
			generateKey: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			if tt.generateKey {
				tt.args.key, err = GenerateKey()
				if err != nil {
					t.Errorf("GenerateKey() error = %v", err)
				}
			}

			encryptText, err := Encrypt(tt.args.plaintext, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := Decrypt(encryptText, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
			}

			if (got != tt.want) != tt.wantErr {
				t.Errorf("got: %v != want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateKey(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "default test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GenerateKey()

			if len(key) < SizeKey {
				t.Errorf("key is invalid error = %v, wantErr %v", err, tt.wantErr)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
