package ui

import (
	"fmt"

	mo "todo-cli/internal/models"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) renderListView() string {
	colorizer := mo.NewColorizer()
	s := colorizer.Colors["purple"] + colorizer.FontStyles["bold"] + "Todo List" + colorizer.Commands["reset"] + "\n\n"

	for i, choice := range m.Todos {

		title := choice.Title
		Cursor := " " // no cursor
		if m.Cursor == i {
			Cursor = colorizer.Colors["pink"] + colorizer.FontStyles["bold"] + ">" + colorizer.Commands["reset"] // cursor!
			title = colorizer.Colors["thickGreen"] + colorizer.FontStyles["bold"] + choice.Title + colorizer.Commands["reset"]
		}

		checked := " " // not selected
		if choice.Done {
			checked = "x" // selected!
		}

		s += fmt.Sprintf("%s [%s] %s\n", Cursor, checked, title)
	}

	return s
}

func (m *Model) updateListView(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		default:
			cmd = m.HandleNavigation(msg)
			m.HandleModification(msg)
			m.HandleOrdering(msg)
		}
	}
	return m, cmd
}
