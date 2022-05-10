package gui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nyan2d/pasm/misc"
)

type CreatePasswordScreen struct {
	items   *GuiBox
	parrent tea.Model
	focus   int
}

func NewCreatePasswordScreen() *CreatePasswordScreen {
	items := NewGuiBox()
	items.Append("title", NewLabel(styleCreatePassTitle.Render("Create new password")))
	items.Append("u_filler", NewFiller(1))
	items.Append("name_label", NewLabel("Name:"))
	items.Append("name_value", NewEditbar())
	items.Append("login_label", NewLabel("Login:"))
	items.Append("login_value", NewEditbar())
	items.Append("pass_label", NewLabel("Password:"))
	items.Append("pass_value", NewEditbar())
	items.Append("info_label", NewLabel("Information:"))
	items.Append("info_value", NewEditbar())
	items.Append("b_filler", NewFiller(0))
	items.Append("help_label", NewLabel("todo: edit this"))

	screen := CreatePasswordScreen{
		items: items,
	}
    // TODO: rewrite
	screen.Init()

	return &screen
}

func (np *CreatePasswordScreen) Init() tea.Cmd {
	np.items.Get("pass_value").(*Editbar).Masked(true)

	cmds := []tea.Cmd{}
	np.items.ForEach(func(x tea.Model) {
		cmds = append(cmds, x.Init())
	})

	return tea.Batch(cmds...)
}

func (np *CreatePasswordScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			// TODO: uncomment
			// return np.parrent, nil
			// TODO: remove
			return np, tea.Quit
		case "tab":
			np.focus = misc.T(np.focus == 3, 0, np.focus+1)
		case "shift+tab":
			np.focus = misc.T(np.focus == 0, 3, np.focus-1)
		}
	case tea.WindowSizeMsg:
		maxHeight := msg.Height
		np.items.Get("b_filler").(*Filler).Height = maxHeight - 11
	}

	// updating childs
	np.items.ForEach(func(x tea.Model) {
		x.Update(msg)
	})

	// updating focus
	np.items.Get("name_value").(*Editbar).Focus(np.focus == 0)
	np.items.Get("login_value").(*Editbar).Focus(np.focus == 1)
	np.items.Get("pass_value").(*Editbar).Focus(np.focus == 2)
	np.items.Get("info_value").(*Editbar).Focus(np.focus == 3)

	return np, nil
}

func (np *CreatePasswordScreen) View() string {
	renderer := NewRenderer()
	np.items.ForEach(func(x tea.Model) {
		renderer.Append(x.View())
	})
	return renderer.Render()
}
