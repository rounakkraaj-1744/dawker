package docker

import (
	"context"
	"fmt"
	"sync"
	"github.com/docker/docker/client"
)

var (
	sharedClient *client.Client
	clientOnce   sync.Once
	clientErr    error
)

func Client() (*client.Client, error) {
	clientOnce.Do(func() {
		sharedClient, clientErr = client.NewClientWithOpts(
			client.FromEnv,
			client.WithAPIVersionNegotiation(),
		)
	})
	if clientErr != nil {
		return nil, fmt.Errorf("docker daemon unavailable: %w", clientErr)
	}
	return sharedClient, nil
}

func NewClient() (*client.Client, error) {
	return Client()
}

func GetContext() context.Context {
	return context.Background()
}