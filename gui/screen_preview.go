package gui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type PreviewScreen struct {
	items *GuiBox
}

func NewPreviewScreen(name, login, pass, info string) *PreviewScreen {
	items := NewGuiBox()
	items.Append("name_caption", NewLabel("Name:"))
	items.Append("name_value", NewLabel(name))
	items.Append("login_caption", NewLabel("Login:"))
	items.Append("login_value", NewLabel(login))
	items.Append("pass_caption", NewLabel("Password:"))
	items.Append("pass_value", NewLabel(pass))
	items.Append("info_caption", NewLabel("Information:"))
	items.Append("info_value", NewLabel(info))
	items.Append("filler", NewFiller(0))
	items.Append("help", NewLabel("privet"))

	return &PreviewScreen{
		items: items,
	}
}

func (ps *PreviewScreen) Init() tea.Cmd {
	cmds := []tea.Cmd{}
	ps.items.ForEach(func(x tea.Model) {
		cmds = append(cmds, x.Init())
	})
	return tea.Batch(cmds...)
}

func (ps *PreviewScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return ps, tea.Quit
		}
	case tea.WindowSizeMsg:
		maxHeight := msg.Height
		ps.items.Get("filler").(*Filler).Height = maxHeight - 9
	}

	ps.items.ForEach(func(x tea.Model) {
		x.Update(msg)
	})
	return ps, nil
}

func (ps *PreviewScreen) View() string {
    renderer := NewRenderer()
	ps.items.ForEach(func(x tea.Model) {
        renderer.Append(x.View())
	})
    return renderer.Render()
}
