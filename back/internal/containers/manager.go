package containers

import (
	"context"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerManager struct {
	cli *client.Client
}

func NewDockerManager() *DockerManager {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error connecting with Docker client: %v", err)
	}

	return &DockerManager{cli: cli}
}

func (dm *DockerManager) ListContainers() ([]container.Summary, error) {
	containers, err := dm.cli.ContainerList(context.Background(), container.ListOptions{All: true})

	return containers, err
}
