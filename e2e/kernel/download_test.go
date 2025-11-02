//go:build integration

package kernel

import (
	"context"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestIntegrationDownloadFile(t *testing.T) {
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

		remoteFile := "/remote.txt"
		localFile := filepath.Join(t.TempDir(), "test.txt")

		err = kernel.Download(connection, remoteFile, localFile)
		require.NoError(t, err)

		got, err := storage.Get(storage.GetDirectionAndFilename(localFile))
		require.NoError(t, err)
		require.Equal(t, "text in remote file", strings.Trim(got, "\n"))
	}
}
