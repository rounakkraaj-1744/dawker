package docker

import (
	"context"
	"strings"

	"uldocker/pkg/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)


func FetchContainers(cli *client.Client) ([]types.Container, error) {
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var results []types.Container
	for _, c := range containers {
		name := "N/A"
		if len(c.Names) > 0 {
			name = strings.TrimPrefix(c.Names[0], "/")
		}

		results = append(results, types.Container{
			ID:     c.ID[:12],
			Name:   name,
			Image:  c.Image,
			State:  c.State,
			Status: c.Status,
		})
	}

	return results, nil
}

func StartContainer(cli *client.Client, id string) error {
	return cli.ContainerStart(context.Background(), id, container.StartOptions{})
}

func StopContainer(cli *client.Client, id string) error {
	return cli.ContainerStop(context.Background(), id, container.StopOptions{})
}