package todo_ui

import (
	"sort"
	cm "todo-cli/internal/enums/command"
	op "todo-cli/internal/enums/operation"
	td "todo-cli/internal/models/todo"
	bh "todo-cli/internal/ui/base"

	"github.com/charmbracelet/bubbletea"
)

func (m *Model) updateAddView(msg tea.Msg) cm.Command {
	todo := m.ShownTodos[m.Cursor]
	switch msg := msg.(type) {
	case tea.KeyMsg: // detects key press
		switch msg.String() {

		case "ctrl+c":
			m.ActiveOp = op.Lister
			m.ShownTodos = append(m.ShownTodos[:m.Cursor], m.ShownTodos[m.Cursor+1:]...)
			return cm.Quit

		case "enter":
			m.Repo.Todos.OrderAndAdd(&m.ShownTodos)
			sort.Sort(td.Todos(m.AllTodos))
			m.ActiveOp = op.Lister
		case "esc":
			m.ActiveOp = op.Lister
			m.ShownTodos = append(m.ShownTodos[:m.Cursor], m.ShownTodos[m.Cursor+1:]...)
			m.Cursor--

		default:
			bh.HandleTextInput(msg, &todo.Title)
			m.ShownTodos[m.Cursor] = todo
		}
	}

	return cm.None
}

func (m *Model) updateEditView(msg tea.Msg) cm.Command {
	todo := m.ShownTodos[m.Cursor]
	switch msg := msg.(type) {
	case tea.KeyMsg: // detects key press
		switch msg.String() {

		case "ctrl+c":
			m.ActiveOp = op.Lister
			return cm.Quit

		case "enter":
			m.Repo.Todos.UpdateField(todo, "Title")
			m.ActiveOp = op.Lister

		case "esc":
			m.ActiveOp = op.Lister
			m.ShownTodos[m.Cursor].Title = m.EditCache

		default:
			bh.HandleTextInput(msg, &todo.Title)
			m.ShownTodos[m.Cursor] = todo

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
			cmd = bh.HandleNavigation(msg, &m.Cursor, len(m.ShownTodos))
			if cmd == cm.None {
				cmd = m.HandleNavUp(msg)
			}
			m.HandleModification(msg)
			m.HandleOrdering(msg)
		}
	}
	return cmd
}
