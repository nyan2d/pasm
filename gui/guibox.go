package gui

import tea "github.com/charmbracelet/bubbletea"

type GuiBox struct {
	items map[string]tea.Model
	names []string
}

func NewGuiBox() *GuiBox {
	return &GuiBox{
		items: map[string]tea.Model{},
		names: []string{},
	}
}

func (gb *GuiBox) Append(name string, model tea.Model) {
	if _, ok := gb.items[name]; !ok {
		gb.names = append(gb.names, name)
	}
	gb.items[name] = model
}

func (gb *GuiBox) Get(name string) tea.Model {
	return gb.items[name]
}

func (gb *GuiBox) ForEach(f func(m tea.Model)) {
	for _, v := range gb.names {
		f(gb.items[v])
	}
}
