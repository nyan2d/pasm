package gui

import "github.com/charmbracelet/lipgloss"

// widget_editbar
var (
	styleEditbarPrompt      = lipgloss.NewStyle().Foreground(lipgloss.Color("#cb00ff"))
	styleEditbarPlaceholder = lipgloss.NewStyle().Foreground(lipgloss.Color("#787878"))
)

// widget_label
var (
	styleLabelText = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#cb00ff"))
)

// widget_list
var (
	styleListCursorActive   = lipgloss.NewStyle().Foreground(lipgloss.Color("#cb00ff"))
	styleListCursorInactive = lipgloss.NewStyle().Foreground(lipgloss.Color("#787878"))
)

// screen_createpass
var (
    styleCreatePassTitle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#cb00ff"))
)
