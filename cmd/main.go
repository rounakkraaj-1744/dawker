package main

import (
	"fmt"
	"os"

	"uldocker/internal/app"
	"uldocker/internal/docker"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cli, err := docker.NewClient()
	if err != nil {
		fmt.Printf("Error initializing Docker client: %v\n", err)
		os.Exit(1)
	}
	defer cli.Close()

	p := tea.NewProgram(app.NewModel(cli), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
