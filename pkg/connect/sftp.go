package connect

import "github.com/pkg/sftp"

type Sftp struct{}

func NewSftp(connection *Connect, opts ...sftp.ClientOption) (*sftp.Client, error) {
	ssh := Ssh{}

	client, err := ssh.Client(connection)
	if err != nil {
		return nil, err
	}

	return sftp.NewClient(client, opts...)
}
