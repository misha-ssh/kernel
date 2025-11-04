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
	require.NoError(t, err)

	defer func() {
		require.NoError(t, c.Terminate(ctx))
	}()

	host, err := c.Host(ctx)
	require.NoError(t, err)

	port, err := c.MappedPort(ctx, "22/tcp")
	require.NoError(t, err)

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
		require.NoError(t, err)

		defer func(session *ssh.Session) {
			err := session.Close()
			if err != nil {
				require.NoError(t, err)
			}
		}(session)

		require.NoError(t, session.Shell())
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
	require.NoError(t, err)

	defer func() {
		require.NoError(t, c.Terminate(ctx))
	}()

	host, err := c.Host(ctx)
	require.NoError(t, err)

	port, err := c.MappedPort(ctx, "22/tcp")
	require.NoError(t, err)

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

// todo add e2e test
func TestIntegrationPrivateKeyConnectWithPassphare(t *testing.T) {}

// todo add e2e test
func TestIntegrationOnlyUsernameAndLoginConnect(t *testing.T) {}

// todo add e2e test
func TestIntegrationOnlyUsernameAndLoginConnectWithPassphare(t *testing.T) {}
