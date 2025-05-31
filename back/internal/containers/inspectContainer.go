package containers

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/josefernando7/node-abyssia/pkg/models/dtos"
)

func InspectContainer(containerId string) (dtos.ContainerInspectDTO, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.45"))
	if err != nil {
		return dtos.ContainerInspectDTO{}, err
	}
	defer cli.Close()

	containerJSON, err := cli.ContainerInspect(context.Background(), containerId)
	if err != nil {
		return dtos.ContainerInspectDTO{}, err
	}

	// Extract ports
	var ports []string
	
	for port, bindings := range containerJSON.NetworkSettings.Ports {
		for _, binding := range bindings {
			ports = append(ports, fmt.Sprintf("%s:%s->%s", binding.HostIP, binding.HostPort, port.Port()))
		}
	}

	// Extract volumes
	var volumes []string

	for _, volume := range containerJSON.Mounts {
		volumes = append(volumes, fmt.Sprintf("%s:%s", volume.Name, volume.Destination))
	}

	return dtos.ContainerInspectDTO{
		ID:      containerJSON.ID,
		Name:    containerJSON.Name,
		Status:  containerJSON.State.Status,
		Image:   containerJSON.Config.Image,
		Ports:   ports,
		Mounts:  volumes,
		StartedAt: containerJSON.State.StartedAt,
	}, nil
}