package gui

import tea "github.com/charmbracelet/bubbletea"

type Gui struct {
}

func NewGui() *Gui {
    return &Gui{}
}

func (g *Gui) Start() error {
    p := tea.NewProgram(NewMainScreen(), tea.WithAltScreen())
    return p.Start()
}
