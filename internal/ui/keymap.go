package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	Quit   key.Binding
	Cmd    key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "j"),
			key.WithHelp("j/up", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "k"),
			key.WithHelp("k/down", "move down"),
		),
		Select: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		Cmd: key.NewBinding(
			key.WithKeys(":"),
			key.WithHelp(":", "enter command mode"),
		),
	}
}