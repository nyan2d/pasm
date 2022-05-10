package gui

import tea "github.com/charmbracelet/bubbletea"

type ItemSelectedMsg struct {
}

func ItemSelectedCmd() tea.Msg {
	return ItemSelectedMsg{}
}
