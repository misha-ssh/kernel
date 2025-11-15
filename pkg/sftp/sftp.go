package sftp

import (
	"github.com/misha-ssh/kernel/pkg/connect"
	ssh2 "github.com/misha-ssh/kernel/pkg/ssh"
	"github.com/pkg/sftp"
)

// Sftp todo put in sftp pkg
type Sftp struct {
	Connection *connect.Connect
}

func (s Sftp) Client(opts ...sftp.ClientOption) (*sftp.Client, error) {
	ssh := &ssh2.Ssh{
		Connection: s.Connection,
	}

	client, err := ssh.Client()
	if err != nil {
		return nil, err
	}

	return sftp.NewClient(client, opts...)
}
