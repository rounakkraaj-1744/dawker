package app

import (
	"uldocker/internal/command"
	"uldocker/internal/state"
	"uldocker/internal/ui"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/docker/docker/client"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	State    *state.AppState
	Docker   *client.Client
	Executor *command.Executor
	KeyMap   ui.KeyMap
	Input    textinput.Model
	Width    int
	Height   int
	Error    error
}

func NewModel(cli *client.Client) Model {
	ti := textinput.New()
	ti.Placeholder = "Enter command..."
	ti.Prompt = ":"
	ti.CharLimit = 156
	ti.Width = 40

	return Model{
		State:    state.NewAppState(),
		Docker:   cli,
		Executor: command.NewExecutor(cli),
		KeyMap:   ui.DefaultKeyMap(),
		Input:    ti,
	}
}

func (m Model) Init() tea.Cmd {
	return FetchContainersCmd(m.Docker)
}