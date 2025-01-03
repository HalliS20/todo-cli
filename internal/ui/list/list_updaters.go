package list_ui

import (
	"sort"
	cm "todo-cli/internal/enums/command"
	op "todo-cli/internal/enums/operation"
	ls "todo-cli/internal/models/lists"
	bh "todo-cli/internal/ui/base"

	"github.com/charmbracelet/bubbletea"
)

func (m *Model) updateAddView(msg tea.Msg) cm.Command {
	todo := m.Lists[m.Cursor]
	switch msg := msg.(type) {
	case tea.KeyMsg: // detects key press
		switch msg.String() {

		case "ctrl+c":
			m.ActiveOp = op.Lister
			m.Lists = append(m.Lists[:m.Cursor], m.Lists[m.Cursor+1:]...)
			return cm.Quit

		case "enter":
			m.Repo.Lists.OrderAndAdd(&m.Lists)
			sort.Sort(ls.Lists(m.Lists))
			m.ActiveOp = op.Lister

		case "esc":
			m.ActiveOp = op.Lister
			m.Lists = append(m.Lists[:m.Cursor], m.Lists[m.Cursor+1:]...)
			m.Cursor--

		default:
			bh.HandleTextInput(msg, &todo.Title)
			m.Lists[m.Cursor] = todo
		}
	}

	return cm.None
}

func (m *Model) updateEditView(msg tea.Msg) cm.Command {
	list := m.Lists[m.Cursor]
	switch msg := msg.(type) {
	case tea.KeyMsg: // detects key press
		switch msg.String() {

		case "ctrl+c":
			m.ActiveOp = op.Lister
			return cm.Quit

		case "enter":
			m.Repo.Lists.UpdateField(list, "Title")
			m.ActiveOp = op.Lister

		case "esc":
			m.ActiveOp = op.Lister
			m.Lists[m.Cursor].Title = m.EditCache

		default:
			bh.HandleTextInput(msg, &list.Title)
			m.Lists[m.Cursor] = list

		}
	}

	return cm.None
}

func (m *Model) updateListView(msg tea.Msg) cm.Command {
	cmd := cm.None
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		default:
			cmd = bh.HandleNavigation(msg, &m.Cursor, len(m.Lists))
			if cmd == cm.None {
				cmd = m.HandleEnteringList(msg)
			}
			m.HandleModification(msg)
			m.HandleOrdering(msg)
		}
	}
	return cmd
}
