package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

func HuhKeyMap() *huh.KeyMap {
	// https://github.com/charmbracelet/huh/blob/main/keymap.go
	return &huh.KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
		),
		Input: huh.InputKeyMap{
			AcceptSuggestion: key.NewBinding(
				key.WithKeys("tab"),
				key.WithHelp("tab", "complete"),
			),
			Prev: key.NewBinding(
				key.WithKeys("shift+tab"),
				key.WithHelp("shift+tab", "back"),
			),
			Next: key.NewBinding(
				key.WithKeys("enter", "tab"),
				key.WithHelp("enter", "next"),
			),
			Submit: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "submit"),
			),
		},
	}
}
