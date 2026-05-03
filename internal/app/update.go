package app

import (
	"uldocker/internal/command"
	"uldocker/internal/docker"
	"uldocker/pkg/types"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/docker/docker/client"
)

func FetchContainersCmd(cli *client.Client) tea.Cmd {
	return func() tea.Msg {
		containers, err := docker.FetchContainers(cli)
		if err != nil {
			return err
		}
		return containers
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil

	case []types.Container:
		m.State.Containers = msg
		return m, nil

	case error:
		m.Error = msg
		return m, nil

	case tea.KeyMsg:
		if m.State.IsCommandMode {
			switch msg.String() {
			case "enter":
				rawCmd := m.Input.Value()
				m.State.IsCommandMode = false
				m.Input.Blur()
				m.Input.Reset()

				if rawCmd != "" {
					parsed := command.Parse(":" + rawCmd)
					if parsed != nil {
						err := m.Executor.Execute(parsed)
						if err != nil {
							m.Error = err
						}
						return m, FetchContainersCmd(m.Docker)
					}
				}
				return m, nil

			case "esc":
				m.State.IsCommandMode = false
				m.Input.Blur()
				m.Input.Reset()
				return m, nil
			}

			m.Input, cmd = m.Input.Update(msg)
			return m, cmd
		}

		switch {
		case key.Matches(msg, m.KeyMap.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.KeyMap.Up):
			if m.State.SelectedIndex > 0 {
				m.State.SelectedIndex--
			}

		case key.Matches(msg, m.KeyMap.Down):
			if m.State.SelectedIndex < len(m.State.Containers)-1 {
				m.State.SelectedIndex++
			}

		case key.Matches(msg, m.KeyMap.Cmd):
			m.State.IsCommandMode = true
			m.Input.Focus()
			return m, nil
		}
	}

	return m, nil
}