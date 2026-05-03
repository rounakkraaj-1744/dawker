package command

import (
	"fmt"
)

func Execute(cmd Command) error {
	switch cmd.Name {

	case "stop":
		fmt.Println("Stopping:", cmd.Args)

	case "start":
		fmt.Println("Starting:", cmd.Args)

	case "restart":
		fmt.Println("Restarting:", cmd.Args)

	default:
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}

	return nil
}