package containers

import (
	"context"

	"github.com/containerd/errdefs"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func DeleteContainer(containerID string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.45"))
	if err != nil {
		return err
	}

	// Check if the container exists
	_, err = cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		if errdefs.IsNotFound(err) {
			return err // Container does not exist, nothing to delete
		}
	}

	// If the container exists, attempt to stop it first
	err = cli.ContainerStop(context.Background(), containerID, container.StopOptions{})
	if err != nil {
		return err
	}

	// Now remove the container
	err = cli.ContainerRemove(context.Background(), containerID, container.RemoveOptions{
		Force: true,
	})
	if err != nil {
		return err
	}

	return nil
}
