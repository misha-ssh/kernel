//go:build integration

package kernel

import (
	"context"
	"os/exec"
	"testing"
	"time"

	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"golang.org/x/crypto/ssh"
)

func TestIntegrationDefaultConnect(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../build/ssh/default",
			Dockerfile: "Dockerfile",
		},
		ExposedPorts: []string{"22/tcp"},
		WaitingFor:   wait.ForListeningPort("22/tcp").WithStartupTimeout(30 * time.Second),
	}

	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := c.Terminate(ctx); err != nil {
			t.Logf("failed to terminate container: %s", err)
		}
	}()

	host, err := c.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}

	port, err := c.MappedPort(ctx, "22/tcp")
	if err != nil {
		t.Fatal(err)
	}

	if c.IsRunning() {
		connection := &connect.Connect{
			Alias:     "test",
			Login:     "root",
			Password:  "password",
			Address:   host,
			Type:      connect.TypeSSH,
			CreatedAt: "",
			UpdatedAt: "",
			SshOptions: &connect.SshOptions{
				Port:       port.Int(),
				PrivateKey: "",
			},
		}

		sshConnector := &connect.Ssh{}
		session, err := sshConnector.Session(connection)
		if err != nil {
			t.Fatal(err)
		}

		defer func(session *ssh.Session) {
			err := session.Close()
			if err != nil {
				t.Fatal(err)
			}
		}(session)

		err = session.Shell()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestIntegrationPrivateKeyConnect(t *testing.T) {
	sshKeysName := "../../build/ssh/key/dockerkey"

	if !storage.Exists(storage.GetDirectionAndFilename(sshKeysName)) {
		cmdKey := exec.Command("ssh-keygen", "-b", "4096", "-t", "rsa", "-f", sshKeysName)
		require.NoError(t, cmdKey.Run())
	}

	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../build/ssh/key",
			Dockerfile: "Dockerfile",
		},
		ExposedPorts: []string{"22/tcp"},
		WaitingFor:   wait.ForListeningPort("22/tcp").WithStartupTimeout(30 * time.Second),
	}

	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := c.Terminate(ctx); err != nil {
			t.Logf("failed to terminate container: %s", err)
		}
	}()

	host, err := c.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}

	port, err := c.MappedPort(ctx, "22/tcp")
	if err != nil {
		t.Fatal(err)
	}

	if c.IsRunning() {
		connection := &connect.Connect{
			Alias:     "test",
			Login:     "root",
			Password:  "password",
			Address:   host,
			Type:      connect.TypeSSH,
			CreatedAt: "",
			UpdatedAt: "",
			SshOptions: &connect.SshOptions{
				Port:       port.Int(),
				PrivateKey: sshKeysName,
			},
		}

		sshConnector := &connect.Ssh{}
		session, err := sshConnector.Session(connection)
		require.NoError(t, err)

		defer func(session *ssh.Session) {
			require.NoError(t, session.Close())
		}(session)

		require.NoError(t, session.Shell())
	}
}
