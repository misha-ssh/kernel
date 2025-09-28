//go:build integration

package kernel

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/misha-ssh/kernel/internal/storage"
	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
	"github.com/misha-ssh/kernel/testutil"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestIntegrationUploadFile(t *testing.T) {
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

		localFile, err := testutil.CreatePrivateKey(t.TempDir())
		require.NoError(t, err)

		remoteFile := "/upload"

		err = kernel.Upload(connection, localFile, remoteFile)
		require.NoError(t, err)

		downloadedFile := filepath.Join(t.TempDir(), "test.txt")

		err = kernel.Download(connection, remoteFile, downloadedFile)
		require.NoError(t, err)

		want, err := storage.Get(storage.GetDirectionAndFilename(localFile))
		require.NoError(t, err)

		got, err := storage.Get(storage.GetDirectionAndFilename(downloadedFile))
		require.NoError(t, err)

		require.Equal(t, want, got)
	}
}
