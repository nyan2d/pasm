package gui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nyan2d/pasm/misc"
)

type Editbar struct {
	text        []rune
	focused     bool
	title       string
	placeholder string
	masked      bool
	onkeyenter  func(text string)
}

func NewEditbar() *Editbar {
	return &Editbar{}
}

func (eb *Editbar) Title(title string) {
	eb.title = title
}

func (eb *Editbar) SetOnKeyEnter(f func(text string)) {
    eb.onkeyenter = f
}

func (eb *Editbar) Placeholder(placeholder string) {
	eb.placeholder = placeholder
}

func (eb *Editbar) Masked(masked bool) {
	eb.masked = masked
}

func (tb *Editbar) Focus(f bool) {
	tb.focused = f
}

func (tb *Editbar) IsFocused() bool {
	return tb.focused
}

func (tb *Editbar) Init() tea.Cmd {
	return nil
}

func (tb *Editbar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !tb.focused {
		return tb, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyBackspace:
			if len(tb.text) > 0 {
				tb.text = tb.text[:len(tb.text)-1]
			}
		case tea.KeyEnter:
            if tb.onkeyenter != nil {
                tb.onkeyenter(string(tb.text))
            }
		case tea.KeyRunes:
			tb.text = append(tb.text, msg.Runes...)
		}
	}

	return tb, nil
}

func (tb *Editbar) View() string {
	prompt := misc.T(tb.focused, styleEditbarPrompt.Render("> "), "  ")
	placeholder := misc.T(len(tb.text) == 0, styleEditbarPlaceholder.Render(tb.placeholder), "")
	content := misc.T(tb.masked, strings.Repeat("*", len(tb.text)), string(tb.text))
	return prompt + placeholder + content
}
