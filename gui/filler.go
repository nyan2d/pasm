package gui

import (
	// "strings"

	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Filler struct {
	Height int
}

func NewFiller(height int) *Filler {
	return &Filler{
		Height: height,
	}
}

func (f *Filler) Init() tea.Cmd {
	return nil
}

func (f *Filler) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return f, nil
}

func (f *Filler) View() string {
	if f.Height > 1 {
		return strings.Repeat("\n", f.Height-1)
	}
	return ""
}
