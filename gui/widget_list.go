package gui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nyan2d/pasm/misc"
)

type ListItem struct {
	ID   int
	Text string
}

type List struct {
	size              *Size
	focused           bool
	items             []string
	filteredItems     []ListItem
	filter            string
	selectedElementId int
	topElementId      int
}

func NewList(items []string, size *Size) *List {
	list := &List{
		items: items,
		size:  size,
	}
	list.bindFilteredItems()

	return list
}

func (lt *List) Focus(f bool) {
	lt.focused = f
}

func (lt *List) SetWidth(width int) {
	lt.size.Width = width
}

func (lt *List) SetHeight(height int) {
	lt.size.Height = height
}

func (lt *List) SetFilter(filter string) {
	lt.filter = filter
	lt.bindFilteredItems()
	lt.selectedElementId = 0
	lt.recalcTopItem()
}

func (lt *List) SetItems(items []string) {
	lt.items = items
	lt.bindFilteredItems()
}

func (lt *List) GetSelectedID() int {
    if len(lt.filteredItems) < 1 {
        return -1
    }

    return lt.filteredItems[lt.selectedElementId].ID
}

func (lt *List) Init() tea.Cmd {
	return nil
}

func (lt *List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if !lt.focused {
			break
		}
		switch msg.String() {
		case "j":
			if lt.selectedElementId < len(lt.filteredItems)-1 {
				lt.selectedElementId++
				lt.recalcTopItem()
			}
		case "k":
			if lt.selectedElementId > 0 {
				lt.selectedElementId--
				lt.recalcTopItem()
			}
		case "enter":
			// return lt, ItemSelectedCmd
			return lt, ItemSelectedCmd

		}
	}
	return lt, nil
}

func (lt *List) View() string {
	return lt.renderItems()
}

func (lt *List) renderItems() string {
	items := []string{}
	cursorActive := styleListCursorActive.Render("# ")
	cursorInactive := styleListCursorInactive.Render("# ")
	cursorInvisible := "  "

	rendered := 0
	for k, v := range lt.filteredItems {
		if k < lt.topElementId {
			continue
		}
		if rendered >= lt.size.Height {
			break
		}

		var builder strings.Builder

		builder.WriteString(misc.T(lt.selectedElementId == k,
			misc.T(lt.focused, cursorActive, cursorInactive), cursorInvisible))

		if lt.size.Width < 1 || len(v.Text) <= lt.size.Width-4 {
			builder.WriteString(v.Text)
		} else {
			builder.WriteString(v.Text[:lt.size.Width-4])
		}

		items = append(items, builder.String())
		rendered++
	}

	return strings.Join(items, "\n")
}

func (lt *List) recalcTopItem() {
	if lt.selectedElementId > lt.topElementId+lt.size.Height-1 {
		lt.topElementId = lt.selectedElementId - lt.size.Height + 1
	}
	if lt.selectedElementId < lt.topElementId && lt.selectedElementId >= 0 {
		lt.topElementId = lt.selectedElementId
	}
}

func (lt *List) bindFilteredItems() {
	if len(lt.filter) <= 0 {
		lt.filteredItems = make([]ListItem, len(lt.items))
		for k, v := range lt.items {
			lt.filteredItems[k] = ListItem{
				ID:   k,
				Text: v,
			}
		}
		return
	}

	lt.filteredItems = make([]ListItem, 0)
	for k, v := range lt.items {
		if misc.IsFuzzyMatch(lt.filter, v) {
			lt.filteredItems = append(lt.filteredItems, ListItem{
				ID:   k,
				Text: v,
			})
		}
	}

}
