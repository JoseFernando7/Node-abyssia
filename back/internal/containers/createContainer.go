package containers

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/josefernando7/node-abyssia/pkg/models"
)

func CreateContainer(request models.ContainerCreateRequest) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.45"))
	if err != nil {
		return "", err
	}

	portBindings := nat.PortMap{}
	exposedPorts := nat.PortSet{}

	for _, port := range request.Ports {
		portParts := nat.Port(port[strings.Index(port, ":")+1:] + "/tcp")
		exposedPorts[portParts] = struct{}{}
		portBindings[portParts] = []nat.PortBinding{{
			HostPort: port[:strings.Index(port, ":")],
		}}
	}

	response, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: 					request.Image,
		ExposedPorts: 	exposedPorts,
	}, &container.HostConfig{
		PortBindings: portBindings,
	}, &network.NetworkingConfig{}, nil, request.Name)

	if err != nil {
		return "", err
	}

	if err := cli.ContainerStart(context.Background(), response.ID, container.StartOptions{}); err != nil {
		return "", err
	}

	return response.ID, nil
}
