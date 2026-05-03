package app

import (
	"fmt"
	"strings"

	"uldocker/internal/ui"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.Error != nil {
		return fmt.Sprintf("Error: %v\n\nPress q to quit.", m.Error)
	}

	var containerList strings.Builder
	containerList.WriteString(ui.HeaderStyle.Render("CONTAINERS") + "\n\n")

	if len(m.State.Containers) == 0 {
		containerList.WriteString("No containers found.")
	} else {
		for i, c := range m.State.Containers {
			cursor := " "
			name := c.Name
			if i == m.State.SelectedIndex {
				cursor = ">"
				name = ui.SelectedStyle.Render(name)
			}
			containerList.WriteString(fmt.Sprintf("%s [%s] %s (%s)\n", cursor, c.ID, name, c.State))
		}
	}

	leftPanel := ui.PanelStyle.
		Width(m.Width/2 - 4).
		Height(m.Height - 6).
		Render(containerList.String())

	rightPanel := ui.PanelStyle.
		Width(m.Width/2 - 4).
		Height(m.Height - 6).
		Render(ui.HeaderStyle.Render("DETAILS / LOGS") + "\n\nLogs will appear here...")

	mainView := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)

	footer := ""
	if m.State.IsCommandMode {
		footer = ui.CommandStyle.Render(m.Input.View())
	} else {
		footer = "j/k: move • enter: select • :: command mode • q: quit"
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		mainView,
		"\n",
		footer,
	)
}