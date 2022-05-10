package gui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Label struct {
	text string
}

func NewLabel(text string) *Label {
	return &Label{
		text: text,
	}
}

func (lb *Label) Text(text string) {
	lb.text = text
}

func (lb *Label) Init() tea.Cmd {
	return nil
}

func (lb *Label) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return lb, nil
}

func (lb *Label) View() string {
    return lb.text
}
