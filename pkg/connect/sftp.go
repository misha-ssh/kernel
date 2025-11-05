package connect

import "github.com/pkg/sftp"

type Sftp struct {
	Connection *Connect
}

func (s Sftp) Client(opts ...sftp.ClientOption) (*sftp.Client, error) {
	ssh := &Ssh{
		Connection: s.Connection,
	}

	client, err := ssh.Client()
	if err != nil {
		return nil, err
	}

	return sftp.NewClient(client, opts...)
}
