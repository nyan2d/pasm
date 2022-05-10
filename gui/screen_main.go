package gui

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nyan2d/pasm/misc"
)

type MainScreen struct {
	items *GuiBox
	focus int
}

func NewMainScreen() *MainScreen {
	searchbar := NewEditbar()
	passlist := NewList([]string{}, &Size{10, 10})

	searchbar.SetOnKeyEnter(func(t string) {
		passlist.SetFilter(t)
	})

	items := NewGuiBox()
	items.Append("searchbar", searchbar)
	items.Append("passlist", passlist)
    // TODO:
	// items.Append("filler", NewFiller(0))
	items.Append("help", NewLabel("haha lol"))

	return &MainScreen{
		items: items,
	}
}

func (ms *MainScreen) Init() tea.Cmd {
	ms.items.Get("searchbar").(*Editbar).Placeholder("Type to search...")

	x := []string{}
	for i := 0; i < 1000; i++ {
		x = append(x, generateString(100))
	}
	ms.items.Get("passlist").(*List).SetItems(x)

	cmds := []tea.Cmd{}
	ms.items.ForEach(func(x tea.Model) {
		cmds = append(cmds, x.Init())
	})
	return tea.Batch(cmds...)
}

func (ms *MainScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return ms, tea.Quit
		case "tab":
			ms.focus = misc.T(ms.focus == 1, 0, ms.focus+1)
		case "shift+tab":
			ms.focus = misc.T(ms.focus == 0, 1, ms.focus-1)
        case "c":
            return NewCreatePasswordScreen(), nil
		}

	case tea.WindowSizeMsg:
		maxX, maxY := msg.Width, msg.Height
		ms.items.Get("passlist").(*List).SetWidth(maxX)
		ms.items.Get("passlist").(*List).SetHeight(maxY - 2)
	}

    //TODO: rewrite
	// updating children
	ms.items.ForEach(func(x tea.Model) {
		_, command := x.Update(msg)
		if command != nil {
			switch command().(type) {
			case ItemSelectedMsg:
				help := ms.items.Get("help").(*Label)
				help.Text(strconv.Itoa(ms.items.Get("passlist").(*List).GetSelectedID()))
			}
		}
	})

	// updating focus
	ms.items.Get("searchbar").(*Editbar).Focus(ms.focus == 0)
	ms.items.Get("passlist").(*List).Focus(ms.focus == 1)

	return ms, nil
}

func (ms *MainScreen) View() string {
	renderer := NewRenderer()
	ms.items.ForEach(func(x tea.Model) {
		renderer.Append(x.View())
	})
	return renderer.Render()
}
