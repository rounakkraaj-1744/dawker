package command

import (
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func Parse(input string) *Command {
	input = strings.TrimPrefix(input, ":")
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}

	return &Command{
		Name: parts[0],
		Args: parts[1:],
	}
}