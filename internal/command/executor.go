package command

import (
	"fmt"
	"uldocker/internal/docker"
	"github.com/docker/docker/client"
)

type Executor struct {
	Registry *Registry
	Client   *client.Client
}

func NewExecutor(cli *client.Client) *Executor {
	reg := NewRegistry()
	
	reg.Register("start", func(c *client.Client, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("usage: :start <container>")
		}
		return docker.StartContainer(c, args[0])
	})

	reg.Register("stop", func(c *client.Client, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("usage: :stop <container>")
		}
		return docker.StopContainer(c, args[0])
	})

	return &Executor{
		Registry: reg,
		Client:   cli,
	}
}

func (e *Executor) Execute(cmd *Command) error {
	handler, ok := e.Registry.Get(cmd.Name)
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(e.Client, cmd.Args)
}