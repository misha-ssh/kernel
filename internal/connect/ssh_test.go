package connect

import "testing"

func TestSshConnect_Connect(t *testing.T) {
	type args struct {
		connect *Connect
	}
	tests := []struct {
		args    args
		name    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				connect: &Connect{
					Alias:     "test",
					Login:     "root",
					Password:  "password",
					Address:   "localhost",
					Type:      TypeSSH,
					CreatedAt: "",
					UpdatedAt: "",
					SshOptions: &SshOptions{
						Port:       22,
						PrivateKey: false,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//s := &SshConnect{}
			//
			//if err := s.Connect(tt.args.connect); (err != nil) != tt.wantErr {
			//	t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			//}
		})
	}
}
